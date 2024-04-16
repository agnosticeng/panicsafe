package panicsafe

import (
	"fmt"
	"runtime/debug"
)

type PanicError struct {
	v     interface{}
	stack []byte
}

func NewPanicError(v interface{}, stack []byte) *PanicError {
	return &PanicError{
		v:     v,
		stack: stack,
	}
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("PanicError(%v): %s", e.v, e.stack)
}

func Defered(retErr *error) {
	if r := recover(); r != nil {
		*retErr = NewPanicError(r, debug.Stack())
	}
}

func Func(f func() error) func() error {
	return func() (retErr error) {
		defer Defered(&retErr)
		return f()
	}
}

func Recover(f func() error) (retErr error) {
	defer Defered(&retErr)
	return f()
}

func Recover2[T any](f func() (T, error)) (ret T, retErr error) {
	defer Defered(&retErr)
	return f()
}
