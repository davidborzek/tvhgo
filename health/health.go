package health

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/go-chi/chi"
)

type healthRouter struct {
	tvhc tvheadend.Client
	db   *sql.DB
}

func New(tvhc tvheadend.Client, db *sql.DB) *healthRouter {
	return &healthRouter{
		tvhc: tvhc,
		db:   db,
	}
}

func (h *healthRouter) Handler() http.Handler {
	r := chi.NewRouter()
	livenessChecker := health.NewChecker()

	readinessChecker := health.NewChecker(
		health.WithTimeout(10*time.Second),
		health.WithCheck(health.Check{
			Name:  "database",
			Check: h.db.PingContext,
		}),
		health.WithCheck(health.Check{
			Name:  "tvheadend",
			Check: h.tvheadendCheck,
		}),
	)

	r.Handle("/liveness", health.NewHandler(livenessChecker))
	r.Handle("/readiness", health.NewHandler(readinessChecker))

	return r
}
