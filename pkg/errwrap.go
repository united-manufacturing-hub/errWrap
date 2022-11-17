package errwrap

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Based on https://smyrman.medium.com/writing-constant-errors-with-go-1-13-10c4191617

type ConstError string

func (err ConstError) Error() string {
	return string(err)
}
func (err ConstError) Is(target error) bool {
	ts := target.Error()
	es := string(err)
	return ts == es || strings.HasPrefix(ts, es+": ")
}
func (err ConstError) Wrap(inner error) error {
	return wrapError{msg: string(err), err: inner}
}
func (err ConstError) WithParams(additionalParams map[string]interface{}) error {
	return wrapError{msg: string(err), additionalParams: additionalParams}
}

func (err ConstError) WrappedWithParams(inner error, additionalParams map[string]interface{}) error {
	return wrapError{msg: string(err), err: inner, additionalParams: additionalParams}
}

type wrapError struct {
	err              error
	additionalParams map[string]interface{}
	msg              string
}

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
func (err wrapError) Unwrap() error {
	return err.err
}
func (err wrapError) Is(target error) bool {
	return ConstError(err.msg).Is(target)
}
