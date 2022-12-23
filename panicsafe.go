package panicsafe

import (
	"fmt"
	"runtime/debug"

	"github.com/go-logr/logr"
)

func Func(f func() error) func() error {
	return func() (retErr error) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)

				if ok {
					retErr = err
				} else {
					retErr = fmt.Errorf("%v", r)
				}

				retErr = fmt.Errorf("%w\n%s", retErr, debug.Stack())
			}
		}()

		return f()
	}
}

func LogErrorFunc(logger logr.Logger, msg string, f func() error) func() {
	return FuncWithErrorLog(logger, msg, f)
}

func FuncWithErrorLog(logger logr.Logger, msg string, f func() error) func() {
	var safeFunc = Func(f)

	return func() {
		var err = safeFunc()

		if err != nil {
			logger.Error(err, msg)
		}
	}
}

func FuncWithErrorHandler(
	f func() error,
	errorHandler func(err error),
) func() {
	var safeFunc = Func(f)

	return func() {
		var err = safeFunc()

		if err != nil {
			errorHandler(err)
		}
	}
}
