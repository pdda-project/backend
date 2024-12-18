package errors

import (
	"github.com/pdda-project/backend/libs/responses"
)

type HTTPError struct {
	Code     int
	Source   error
	Response responses.Error
}

func (e *HTTPError) Error() string {
	if e.Source != nil {
		return e.Source.Error()
	} else {
		return "HTTPError"
	}
}

func NewHTTPError(code int, errors []string, source error) *HTTPError {
	return &HTTPError{
		Code:   code,
		Source: source,
		Response: responses.Error{
			Status: "fail",
			Errors: errors,
		},
	}
}
