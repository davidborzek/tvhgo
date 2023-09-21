package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed dist
var dist embed.FS

func NewRouter() (http.Handler, error) {
	r := chi.NewRouter()

	uiFS, err := fs.Sub(dist, "dist")
	if err != nil {
		return nil, err
	}

	srv := http.FileServer(http.FS(uiFS))
	r.Get("/", srv.ServeHTTP)
	r.Get("/manifest.webmanifest", srv.ServeHTTP)
	r.Get("/favicon.ico", srv.ServeHTTP)
	r.Get("/assets/*", srv.ServeHTTP)
	r.Get("/img/*", srv.ServeHTTP)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		srv.ServeHTTP(w, r)
	})

	return r, nil
}
