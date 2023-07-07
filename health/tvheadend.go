package health

import (
	"context"
	"fmt"
)

func (r *healthRouter) tvheadendCheck(ctx context.Context) error {
	res, err := r.tvhc.Exec(ctx, "/status.xml", nil)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("tvheadend returned erroneous status code: %d", res.StatusCode)
	}

	return nil
}
