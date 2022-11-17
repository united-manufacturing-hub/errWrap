# errWrap

errWrap is a simple package to allow constant error handling in Go & error wrapping. 

Expect breaking changes until v1.0.0.


## Contributing

Make sure to install the pre-push hooks.

    $ git config core.hooksPath .githooks

## Importing

   go get github.com/united-manufacturing-hub/errWrap

## Usage

 - Define your errors as `const YourErr = ConstError("error message")`
   - You can wrap errors with `YourError.Wrap(error)`
   - You can add additional context with `YourError.WithParams(map[string]interface{})`
   - Combine the two with `YourError.WrappedWithParams(error,map[string]interface{})`