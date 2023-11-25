package response

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	ErrorResponse struct {
		Message string `json:"message"`
	}
)

// Error writes the json encoded error message to the response.
func Error(w http.ResponseWriter, err error, status int) {
	JSON(w, &ErrorResponse{Message: err.Error()}, status)
}

// InternalError writes the json encoded error message to the response
// with 500 internal server error status code.
func InternalError(w http.ResponseWriter, err error) {
	Error(w, err, 500)
}

// InternalError writes the json encoded error message to the response
// with 500 internal server error status code.
func InternalErrorCommon(w http.ResponseWriter) {
	InternalError(w, errors.New("unexpected error"))
}

// Unauthorized writes the json encoded error message to the response
// with 401 unauthorized error status code.
func Unauthorized(w http.ResponseWriter, err error) {
	Error(w, err, 401)
}

// Unauthorized writes the json encoded error message to the response
// with 401 unauthorized error status code.
func Unauthorizedf(w http.ResponseWriter, format string, a ...any) {
	Error(w, fmt.Errorf(format, a...), 401)
}

// BadRequest writes the json encoded error message to the response
// with 400 bad request status code.
func BadRequest(w http.ResponseWriter, err error) {
	Error(w, err, 400)
}

// BadRequest writes the json encoded error message to the response
// with 400 bad request status code.
func BadRequestf(w http.ResponseWriter, format string, a ...any) {
	BadRequest(w, fmt.Errorf(format, a...))
}

// NotFound writes the json encoded error message to the response
// with 404 not found status code.
func NotFound(w http.ResponseWriter, err error) {
	Error(w, err, 404)
}

// Conflict writes the json encoded error message to the response
// with 409 conflict status code.
func Conflict(w http.ResponseWriter, err error) {
	Error(w, err, 409)
}

// Forbidden writes the json encoded error message to the response
// with 403 forbidden error status code.
func Forbidden(w http.ResponseWriter, err error) {
	Error(w, err, 403)
}

// Forbiddenf writes the json encoded error message to the response
// with 403 forbidden error status code.
func Forbiddenf(w http.ResponseWriter, format string, a ...any) {
	Error(w, fmt.Errorf(format, a...), 403)
}
