package errors

import (
	"encoding/json"
	"net/http"

	"github.com/pdda-project/backend/utils/responses"
)

var (
	InternalError     = NewHTTPError(http.StatusInternalServerError, "internal server error")
	ForbiddenError    = NewHTTPError(http.StatusForbidden, "forbidden")
	UnauthorizedError = NewHTTPError(http.StatusUnauthorized, "unauthorized")
)

type HTTPError struct {
	Code    int
	Message string
}

func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

func (e *HTTPError) Error() string {
	return e.Message
}

func (e *HTTPError) Send(w http.ResponseWriter) {
	w.WriteHeader(e.Code)
	res := &responses.ResponseMessage{
		Status:  "fail",
		Message: e.Message,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		res := map[string]string{
			"status":  "fail",
			"message": "json encoder error",
		}
		json.NewEncoder(w).Encode(res)
	}
}
