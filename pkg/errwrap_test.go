package pkg

import (
	"testing"
)

const exampleError = ConstError("example error")
const wrappedError = ConstError("wrapped error")

func Test_ConstError(t *testing.T) {
	err := exampleError
	if err.Error() != "example error" {
		t.Errorf("Error() = %v, want %v", err.Error(), "example error")
	}
}

func Test_ConstErrorWrapped(t *testing.T) {
	err := exampleError.Wrap(wrappedError)
	if err.Error() != "example error: wrapped error" {
		t.Errorf("Error() = %v, want %v", err.Error(), "example error: wrapped error")
	}
}

func Test_ConstErrorParams(t *testing.T) {
	params := map[string]any{"param1": "value1", "param2": "value2"}
	err := exampleError.WithParams(params)
	if err.Error() != "example error [{\"param1\":\"value1\",\"param2\":\"value2\"}]" {
		t.Errorf("Error() = %v, want %v", err.Error(), "example error [{\"param1\":\"value1\",\"param2\":\"value2\"}]")
	}
}

func TestConstErrorWrappedParams(t *testing.T) {
	params := map[string]any{"param1": "value1", "param2": "value2"}
	err := exampleError.WrappedWithParams(wrappedError, params)
	if err.Error() != "example error: wrapped error [{\"param1\":\"value1\",\"param2\":\"value2\"}]" {
		t.Errorf(
			"Error() = %v, want %v",
			err.Error(),
			"example error: wrapped error [{\"param1\":\"value1\",\"param2\":\"value2\"}]")
	}
}
