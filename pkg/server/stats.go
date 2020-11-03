package server

import (
	"encoding/json"
	"errors"
	"sync"
)

type StatsRequest interface {
	StatsKey() string
}

type Stats struct {
	Api   string `json:"api"`
	Body  string `json:"body"`
	Count uint   `json:"count"`
}

// StatsRepository is a repository to store some stats for a request
type StatsRepository interface {
	Add(r StatsRequest) error
	MostUsed(skey string) Stats
}

type memoryStats struct {
	requests map[string]map[string]uint
	mu       sync.Mutex
}

func newStatsRepository() StatsRepository {
	return &memoryStats{
		requests: make(map[string]map[string]uint),
	}
}

func (repo *memoryStats) Add(r StatsRequest) error {
	if r == nil {
		return errors.New("nil request")
	}
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	key := r.StatsKey()
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, ok := repo.requests[key]; !ok {
		repo.requests[key] = make(map[string]uint)
	}
	repo.requests[key][string(b)]++
	return nil
}

func (repo *memoryStats) MostUsed(skey string) Stats {
	stats := Stats{
		Api: skey,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()
	bodies, ok := repo.requests[skey]
	if !ok {
		return stats
	}

	max := uint(0)
	selected := ""
	for k, v := range bodies {
		if v > max {
			max = v
			selected = k
		}
	}

	if cnt, ok := bodies[selected]; ok {
		stats.Count = cnt
		stats.Body = selected
	}

	return stats
}
