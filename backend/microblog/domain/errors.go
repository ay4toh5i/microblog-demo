package domain

import "errors"

var (
	ErrInvalidRequest = errors.New("The request is invalid.")
	ErrNotFound       = errors.New("The entity is not found.")
	ErrAuthentication = errors.New("Valid credential required.")
)
