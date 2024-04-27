package api

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func (router *router) Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			log.Debug().Fields(map[string]any{
				"remote_addr": r.RemoteAddr,
				"path":        r.URL.Path,
				"proto":       r.Proto,
				"method":      r.Method,
				"host":        r.Host,
				"user_agent":  r.UserAgent(),
				"status_code": ww.Status(),
			}).Msg("api request")
		}()

		next.ServeHTTP(ww, r)
	}

	return http.HandlerFunc(fn)
}
