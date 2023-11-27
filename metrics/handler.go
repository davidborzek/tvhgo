package metrics

import (
	"errors"
	"net/http"
	"strings"

	"github.com/davidborzek/tvhgo/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// authenticate authenticates a request with a expected token.
func authenticate(r *http.Request, expectedToken string) error {
	token := strings.ReplaceAll(
		r.Header.Get("Authorization"),
		"Bearer ", "")

	if token != expectedToken {
		return errors.New("authentication failed")
	}

	return nil
}

func Handler(cfg config.MetricsConfig, collectors ...prometheus.Collector) func(http.ResponseWriter, *http.Request) {
	if !cfg.Enabled {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}
	}

	for _, collector := range collectors {
		prometheus.MustRegister(collector)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := authenticate(r, cfg.Token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		promhttp.Handler().ServeHTTP(w, r)
	}
}
