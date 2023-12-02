package validator

import "errors"

var (
	ErrInvalidEmail     = errors.New("invalid mail")
	ErrInvalidRange     = errors.New("invalid range")
	ErrrInvalidPassword = errors.New("invalid password")
)
