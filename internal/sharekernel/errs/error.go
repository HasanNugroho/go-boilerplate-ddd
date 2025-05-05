package errs

import (
	"errors"
	"fmt"
)

// Sentinel Errors (konstanta sederhana)
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")
	ErrInternal     = errors.New("internal server error")
	ErrConflict     = errors.New("conflict")
)

// NotFoundError dengan informasi sumber daya dan ID
type NotFoundError struct {
	Resource string
	ID       interface{}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID '%v' not found", e.Resource, e.ID)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound) || errors.As(err, &NotFoundError{})
}

// ValidationError dengan informasi field dan pesan
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

func IsValidationError(err error) bool {
	return errors.As(err, &ValidationError{})
}

// InternalError dengan kode dan error asli
type InternalError struct {
	Code    string
	Message string
	Err     error
}

func (e InternalError) Error() string {
	msg := fmt.Sprintf("internal error (code: %s): %s", e.Code, e.Message)
	if e.Err != nil {
		msg += fmt.Sprintf(" - original error: %v", e.Err)
	}
	return msg
}

func NewInternalError(code, message string, original error) error {
	return InternalError{Code: code, Message: message, Err: original}
}

func IsInternalError(err error) bool {
	return errors.As(err, &InternalError{})
}
