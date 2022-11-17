package errwrap

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Based on https://smyrman.medium.com/writing-constant-errors-with-go-1-13-10c4191617

// ConstError is an error that is constant.
type ConstError string

// Error converts the error into a string.
func (err ConstError) Error() string {
	return string(err)
}

// Is checks if the given error is equal to this error.
func (err ConstError) Is(target error) bool {
	ts := target.Error()
	es := string(err)
	return ts == es || strings.HasPrefix(ts, es+": ")
}

// Wrap wraps the given error into this error.
func (err ConstError) Wrap(inner error) error {
	return wrapError{msg: string(err), err: inner}
}

// WithParams returns an error with additional params.
func (err ConstError) WithParams(additionalParams map[string]interface{}) error {
	return wrapError{msg: string(err), additionalParams: additionalParams}
}

// WrappedWithParams returns an error that wraps the given error, and adds the given additionalParams.
func (err ConstError) WrappedWithParams(inner error, additionalParams map[string]interface{}) error {
	return wrapError{msg: string(err), err: inner, additionalParams: additionalParams}
}

// wrapError is an error that wraps another error.
type wrapError struct {
	err              error
	additionalParams map[string]interface{}
	msg              string
}

// Error formats the error message, and optionally it's inner error & additional params.
func (err wrapError) Error() string {
	if err.err != nil {
		if len(err.additionalParams) > 0 {
			marshal, errM := json.Marshal(err.additionalParams)
			if errM != nil {
				return fmt.Sprintf("%s: %v {FAILED TO MARSHAL ADDITIONAL PARAMS}", err.msg, err.err)
			}
			return fmt.Sprintf("%s: %v [%s]", err.msg, err.err, marshal)
		} else {
			return fmt.Sprintf("%s: %v", err.msg, err.err)
		}
	}
	if len(err.additionalParams) > 0 {
		marshal, errM := json.Marshal(err.additionalParams)
		if errM != nil {
			panic(errM)
		}
		return fmt.Sprintf("%s [%s]", err.msg, marshal)
	}
	return err.msg
}

// Unwrap returns the inner error.
func (err wrapError) Unwrap() error {
	return err.err
}

// Is converts the message into an ConstError and calls ConstError.Is on that.
func (err wrapError) Is(target error) bool {
	return ConstError(err.msg).Is(target)
}
