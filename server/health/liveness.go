package health

import (
	"context"
	"net/http"

	"github.com/davidborzek/tvhgo/api/response"
)

func (h *router) handleLiveness(w http.ResponseWriter, r *http.Request) {
	components := h.liveness(r.Context())

	statusText := statusUp
	status := 200

	for _, c := range components {
		if c.Status == statusDown {
			statusText = statusDown
			status = 503
			break
		}
	}

	response.JSON(w, healthStatus{
		Status:     statusText,
		Components: components,
	}, status)
}

func (h *router) liveness(ctx context.Context) map[string]componentStatus {
	return map[string]componentStatus{}
}
