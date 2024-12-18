package utils

import (
	"encoding/json"
	"net/http"

	"github.com/pdda-project/backend/libs/errors"
)

func FromRequest[T any](request *http.Request, reader *T) *errors.HTTPError {
	var httpErr *errors.HTTPError = nil

	defer request.Body.Close()
	err := json.NewDecoder(request.Body).Decode(reader)
	if err != nil {
		httpErr = errors.NewHTTPError(400, []string{"invalid payload", err.Error()}, err)
	}
	return httpErr
}
