package validation

import "github.com/go-playground/validator"

var (
	validate = validator.New()
)

func ValidateStruct[T any](s T) (T, error) {
	err := validate.Struct(s)
	return s, err
}
