package errorsx

import (
	"errors"
	"fmt"

	"github.com/faisallbhr/light-pos-be/config"
)

type ErrorWithCode struct {
	Code    error
	Message string
	Detail  error
}

func (e *ErrorWithCode) Error() string {
	if config.IsProduction() {
		return e.Message
	}
	return fmt.Sprintf("%s: %s", e.Message, e.Detail)
}

func productionError(code error, msg string) error {
	return &ErrorWithCode{
		Code:    code,
		Message: msg,
	}
}

func NewError(code error, msg string, err error) error {
	if config.IsProduction() {
		return productionError(code, msg)
	}
	return &ErrorWithCode{
		Code:    code,
		Message: msg,
		Detail:  err,
	}
}

func GetCode(err error) error {
	var e *ErrorWithCode
	if errors.As(err, &e) {
		return e.Code
	}
	return ErrInternal
}
