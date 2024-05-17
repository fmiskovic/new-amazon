package repositories

import "errors"

var (
	ErrNilEntity = errors.New("entity can not be nil")
	ErrNotFound  = errors.New("entity not found")
)
