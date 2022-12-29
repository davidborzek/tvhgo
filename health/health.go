package health

import (
	"net/http"

	"github.com/davidborzek/tvhgo/api/response"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type componentStatus struct {
	Status  statusText `json:"status"`
	Message string     `json:"message,omitempty"`
}

type healthStatus struct {
	Status     statusText                 `json:"status"`
	Components map[string]componentStatus `json:"components,omitempty"`
}

type router struct {
	tvhc tvheadend.Client
}

type statusText string

const (
	statusDown statusText = "DOWN"
	statusUp   statusText = "UP"
)

func New(tvhc tvheadend.Client) *router {
	return &router{
		tvhc: tvhc,
	}
}

func (h *router) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Get("/", h.handleHealth)
	r.Get("/liveness", h.handleLiveness)
	r.Get("/readiness", h.handleReadiness)

	return r
}

func (h *router) handleHealth(w http.ResponseWriter, r *http.Request) {
	liveness := h.liveness(r.Context())
	readiness := h.readiness(r.Context())

	components := mergeComponents(liveness, readiness)

	statusText := statusUp
	status := 200

	for _, c := range components {
		if c.Status == statusDown {
			statusText = statusDown
			status = 503
			break
		}
	}

	res := healthStatus{
		Status:     statusText,
		Components: components,
	}

	response.JSON(w, res, status)
}

func mergeComponents(components ...map[string]componentStatus) map[string]componentStatus {
	out := make(map[string]componentStatus)

	for _, c := range components {
		for name, status := range c {
			out[name] = status
		}
	}

	return out
}
