package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Server interface {
	Run() error
}

func New() Server {
	return &server{
		router: chi.NewRouter(),
		stats:  newStatsRepository(),
	}
}

type server struct {
	router chi.Router
	stats  StatsRepository
}

func (s *server) init() error {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(30 * time.Second))
	s.router.Use(WithStats(&apiStatsFormatter{}, s.stats))

	// Index
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
<!Doctype>
<html>
	<head>
		<title>Fizzbuzz - Thomas Charlot</title>
	</head>
	<body>
		<h1>Fizzbuzz !</h1>
		<p>Try <a href="/fizzbuzz">api</a> or <a href="/stats">stats</a></p>
	</body>
</html>`))
	})

	// Fizzbuzz endpoint
	s.router.Post("/fizzbuzz", s.postFizzbuzz)
	s.router.Get("/fizzbuzz", s.getFizzbuzz)

	// Stats endpoint
	s.router.Get("/stats/all", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, s.stats.All())
	})
	s.router.Get("/stats/mostused", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, s.stats.MostUsed())
	})

	return nil
}

func (s *server) Run() error {
	if err := s.init(); err != nil {
		return err
	}

	return http.ListenAndServe(":8081", s.router)
}
