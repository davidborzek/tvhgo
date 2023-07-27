package response

import (
	"encoding/json"
	"io"
	"net/http"
)

// JSON writes the json-encoded data to the response
// with a given status code.
func JSON(w http.ResponseWriter, v interface{}, status int) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return enc.Encode(v)
}

// CopyResponse writes a http response to the response
// by copying the header, status code and the body.
func CopyResponse(dest http.ResponseWriter, src *http.Response) (written int64, err error) {
	defer src.Body.Close()
	copyHeaders(dest.Header(), src.Header)
	dest.WriteHeader(src.StatusCode)
	return io.Copy(dest, src.Body)
}

func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
