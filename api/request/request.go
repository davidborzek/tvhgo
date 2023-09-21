package request

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

var (
	ErrFailedToBindJSON = errors.New("failed to parse json body")
)

// BindJSON decodes the body of a http request to the
// provided struct.
func BindJSON(r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return ErrFailedToBindJSON
	}

	r.Body.Close()
	return nil
}

// BindQuery decodes the url query of a http request to the
// provided struct.
func BindQuery(r *http.Request, dst interface{}) error {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)

	return d.Decode(dst, r.URL.Query())
}

// RemoteAddr returns the "real" IP address.
func RemoteAddr(r *http.Request) string {
	addr := r.Header.Get("X-Real-IP")

	if addr == "" {
		addr = strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
		addr = strings.TrimSpace(addr)
	}

	if addr != "" && net.ParseIP(addr) == nil {
		addr = ""
	}

	if addr == "" {
		addr, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return addr
}

// NumericURLParam returns a url parameter as int64.
func NumericURLParam(r *http.Request, name string) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, name), 10, 0)
}
