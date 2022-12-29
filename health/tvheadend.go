package health

import (
	"context"
	"fmt"
)

func (r *router) tvheadendStatus(ctx context.Context) componentStatus {
	res, err := r.tvhc.Exec(ctx, "/status.xml", nil)
	if err != nil {
		return componentStatus{
			Status:  statusDown,
			Message: err.Error(),
		}
	}

	if res.StatusCode >= 400 {
		return componentStatus{
			Status:  statusDown,
			Message: fmt.Sprintf("tvheadend returned erroneous status code: %d", res.StatusCode),
		}
	}

	return componentStatus{
		Status: statusUp,
	}
}
