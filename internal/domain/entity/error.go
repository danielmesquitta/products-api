package entity

import (
	"fmt"
	"runtime/debug"
)

type Err struct {
	Message    string
	StackTrace string
	Type       ErrType
}

type ErrType string

const (
	ErrTypeUnknown    ErrType = "unknown"
	ErrTypeNotFound   ErrType = "not_found"
	ErrTypeValidation ErrType = "validation_error"
)

func newErr(err any, errType ErrType) *Err {
	switch v := err.(type) {
	case *Err:
		return v
	case error:
		return &Err{
			Message:    v.Error(),
			StackTrace: string(debug.Stack()),
			Type:       errType,
		}
	case string:
		return &Err{
			Message:    v,
			StackTrace: string(debug.Stack()),
			Type:       errType,
		}
	default:
		panic("trying to create an Err with an unsupported type")
	}
}

// NewErr creates a new Err instance from either an error or a string,
// and sets the Type flag to unknown. This is useful when you want to
// create an error that is not expected to happen, and you want to
// log it with stack tracing.
func NewErr(err any) *Err {
	return newErr(err, ErrTypeUnknown)
}

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) ErrorWithStackTrace() string {
	return fmt.Sprintf("%s\n\n%s", e.Message, e.StackTrace)
}

var (
	ErrProductNotFound = newErr("product not found", ErrTypeNotFound)
	ErrValidation      = newErr("validation error", ErrTypeValidation)
)
