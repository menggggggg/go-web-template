package errors

import "github.com/pkg/errors"

var (
	ErrBadRequest          = "BadRequest"
	ErrInternalServerError = "InternalServerError"
)

func WrapWithBadRequest(err error) error {
	return errors.Wrap(err, ErrBadRequest)
}

func WrapWithInternalServerError(err error) error {
	return errors.Wrap(err, ErrInternalServerError)
}
