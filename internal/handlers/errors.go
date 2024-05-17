package handlers

import (
	"fmt"
)

type HandlerError struct {
	message string
	err     error
	code    int
}

func NewErr(msg string, err error, code int) HandlerError {
	return HandlerError{
		message: msg,
		err:     err,
		code:    code,
	}
}

func (h HandlerError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s, error: %v", h.code, h.message, h.err)
}
