package rest

import (
	"fmt"
	"net/http"
)

// Errors is a collection of REST errors
type Errors struct {
	Errors []*Error `json:"errors"`
}

// Error represent an error returned by the REST API
type Error struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("ID: %q, Status: %d, Title: %q, Detail: %q", e.ID, e.Status, e.Title, e.Detail)
}

func writeError(w http.ResponseWriter, r *http.Request, err *Error) {
	w.WriteHeader(err.Status)
	encodeJSONResponse(w, r, Errors{[]*Error{err}})
}

var (
	errNotFound  = &Error{"not_found", 404, "Not Found", "Requested content not found."}
	errForbidden = &Error{"forbidden", 401, "Forbidden", "This operation is forbidden."}
)

func newContentNotFoundError(contentName string) *Error {
	return &Error{"not_found", 404, "Not Found", fmt.Sprintf("%s not found.", contentName)}
}

func newInternalServerError(err interface{}) *Error {
	return &Error{"internal_server_error", 500, "Internal Server Error", fmt.Sprintf("Something went wrong: %+v", err)}
}

func newNotAcceptableError(accept string) *Error {
	return &Error{"not_acceptable", 406, "Not Acceptable", fmt.Sprintf("Accept header must be set to '%s'.", accept)}
}

func newUnsupportedMediaTypeError(contentType string) *Error {
	return &Error{"unsupported_media_type", 415, "Unsupported Media Type", fmt.Sprintf("Content-Type header must be set to: '%s'.", contentType)}
}

func newBadRequestParameter(param string, err error) *Error {
	return &Error{"bad_request", http.StatusBadRequest, "Bad Request", fmt.Sprintf("Invalid %q parameter %v", param, err)}
}

func newBadRequestError(err error) *Error {
	return &Error{"bad_request", http.StatusBadRequest, "Bad Request", fmt.Sprint(err)}
}

func newBadRequestMessage(message string) *Error {
	return &Error{"bad_request", http.StatusBadRequest, "Bad Request", message}
}

func newConflictRequest(message string) *Error {
	return &Error{"conflict", http.StatusConflict, "Conflict", message}
}

func newForbiddenRequest(message string) *Error {
	return &Error{"forbidden", http.StatusForbidden, "Forbidden", message}
}
