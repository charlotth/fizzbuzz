package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
)

type Stats struct {
	Api    string `json:"api"`
	Params string `json:"params"`
	Count  uint   `json:"count"`
}

type StatsRequestFormatter interface {
	NewStatsRequest(r *http.Request) (key string, values map[string]string)
}

// StatsRepository is a repository to store some stats for a request
type StatsRepository interface {
	Add(key string, values map[string]string) error
	MostUsed() *Stats
	All() []*Stats
}

// WithStats is our middleware
func WithStats(f StatsRequestFormatter, repo StatsRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			key, values := f.NewStatsRequest(r)
			repo.Add(key, values)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

type apiStatsFormatter struct {
}

func (f *apiStatsFormatter) NewStatsRequest(r *http.Request) (string, map[string]string) {
	key := fmt.Sprintf("%s %s", r.Method, r.RequestURI)
	values := make(map[string]string)

	// Copy query params
	for k, v := range r.URL.Query() {
		sort.Strings(v)
		values[k] = strings.Join(v, ",")
	}

	// Read body
	body := f.drainBody(r)
	if len(body) > 0 {
		values["body"] = body
	}
	return key, values
}

func (f *apiStatsFormatter) drainBody(r *http.Request) string {
	if r.Body == nil || r.Body == http.NoBody {
		return ""
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return ""
	}
	if err := r.Body.Close(); err != nil {
		return ""
	}
	body := buf.Bytes()
	r.Body = io.NopCloser(bytes.NewReader(body))
	return string(body)
}

type memoryStats struct {
	requests map[string]uint
	mu       sync.Mutex
}

func newStatsRepository() StatsRepository {
	return &memoryStats{
		requests: make(map[string]uint),
	}
}

func (repo *memoryStats) createKey(key string, values map[string]string) (string, error) {
	params := ""

	if len(values) > 0 {
		b, err := json.Marshal(values)
		if err != nil {
			return "", err
		}
		params = string(b)
	}

	return fmt.Sprintf("%s###%s", key, params), nil
}

func (repo *memoryStats) Add(key string, values map[string]string) error {
	rkey, err := repo.createKey(key, values)
	if err != nil {
		return err
	}
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if cnt, ok := repo.requests[rkey]; !ok {
		repo.requests[rkey] = 1
	} else {
		repo.requests[rkey] = cnt + 1
	}
	return nil
}

func (repo *memoryStats) MostUsed() *Stats {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	max := uint(0)
	selected := ""

	for k, v := range repo.requests {
		if v > max {
			max = v
			selected = k
		}
	}

	parts := strings.Split(selected, "###")
	if len(parts) != 2 {
		return nil
	}

	return &Stats{
		Api:    parts[0],
		Params: parts[1],
		Count:  max,
	}
}

func (repo *memoryStats) All() []*Stats {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	stats := make([]*Stats, 0, len(repo.requests))
	for k, v := range repo.requests {
		parts := strings.Split(k, "###")
		if len(parts) != 2 {
			continue
		}
		stats = append(stats, &Stats{
			Api:    parts[0],
			Params: parts[1],
			Count:  v,
		})
	}
	return stats
}
