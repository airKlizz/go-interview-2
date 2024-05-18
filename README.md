# Goome

## 6 - Errors

### Objective

We have seen that when an error occured, we have no information about it when using the CLI or the HTTP server.
We would like to differenciate the different errors, to be able to respond the correct HTTP code for example.

### Errors in Golang

In Golang create a custom error is as simple as:

```go
var ErrNoObjectsFound = errors.New("no objects found")
```

This works in many cases and is widely used.

This approach allows to differentiate errors from each other but does not allow to add context to the error.
For exemple, we can create a custom error for an event not valid:

```go
var ErrEventNotValid = errors.New("event not valid")
```

but we cannot add the reasons why the event is not valid.

To do so, we can create a new struct that implements the `error` interface:

```go
type error interface {
	Error() string
}
```

### ErrEventNotValid

Let's create this custom error.

In the `internal/core/domain` folder, create a new file called `errors.go` with the following content:

```go
package domain

import (
	"fmt"
	"strings"

	cue "cuelang.org/go/cue/errors"
)

var (
	ErrorEventNotValid = ErrEventNotValid{} // event is not valid
)

type ErrEventNotValid struct {
	details []string
}

func NewErrEventNotValid(err error) error {
	details := []string{}
	for _, err := range cue.Errors(err) {
		details = append(details, fmt.Sprintf("%s %s", err.Error(), strings.Join(err.Path(), "/")))
	}
	return ErrEventNotValid{
		details: details,
	}
}

func (err ErrEventNotValid) Error() string {
	return fmt.Sprintf("event not valid: %s", strings.Join(err.details, ", "))
}

func (err ErrEventNotValid) Is(target error) bool {
	_, ok := target.(ErrEventNotValid)
	return ok
}

We create a struct `ErrEventNotValid` with as details the error message coming from the CUE validation.
In the constructor we create the details of the error.
The struct implements the `error` interface with the `Error() string` method.
We also implement the `Is(target error) bool` method to make the `Is` function from the `errors` standard package working (see [doc](https://pkg.go.dev/errors#Is) for details).

### Usage of the custom error

## Next

To see the solutions:

```bash
git checkout 5-data-validation-end
```

To go to the next step:

```bash
git checkout 6-errors
```
