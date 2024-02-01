package metrics

import (
	"net/http"
	"strings"

	"github.com/davidborzek/tvhgo/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

type server struct {
	cfg        *config.MetricsConfig
	collectors []prometheus.Collector
}

func NewServer(cfg *config.MetricsConfig, collectors ...prometheus.Collector) *server {
	return &server{
		cfg:        cfg,
		collectors: collectors,
	}
}

func (s *server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Use(s.authenticate)

	for _, collector := range s.collectors {
		prometheus.MustRegister(collector)
	}

	r.Handle("/metrics", promhttp.Handler())

	addr := s.cfg.Addr()
	log.Info().Str("addr", addr).
		Msg("starting the metrics http server")

	go func() {
		if err := http.ListenAndServe(addr, r); err != nil {
			log.Fatal().Err(err).Msg("failed to start metrics http server")
		}
	}()
}

func (s *server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.ReplaceAll(
			r.Header.Get("Authorization"),
			"Bearer ", "")

		if len(s.cfg.Token) != 0 && token != s.cfg.Token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
