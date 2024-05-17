package services

import "fmt"

type ServiceError struct {
	message string
	err     error
}

func newError(message string, err error) ServiceError {
	return ServiceError{message: message, err: err}
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("messge: %s, error: %v", e.message, e.err)
}
