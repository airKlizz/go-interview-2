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
)

var (
	ErrorEventNotValid = ErrEventNotValid{} // event is not valid
)

type ErrEventNotValid struct {
	reasons []string
}

func NewErrEventNotValid(reasons []string) error {
	return ErrEventNotValid{
		reasons: reasons,
	}
}

func (err ErrEventNotValid) Error() string {
	return fmt.Sprintf("event not valid: %s", strings.Join(err.reasons, ", "))
}

func (err ErrEventNotValid) Is(target error) bool {
	_, ok := target.(ErrEventNotValid)
	return ok
}

We create a struct `ErrEventNotValid` with as reasons a list of string coming from the validation of the event.
The struct implements the `error` interface with the `Error() string` method.
We also implement the `Is(target error) bool` method to make the `Is` function from the `errors` standard package working (see [doc](https://pkg.go.dev/errors#Is) for details).

You can have a look to the `Validate` method of the `Event` struct which has been slighly modified to return the reasons when an event is not valid.

### More custom errors

We have seen how to create a custom error with context stored into the error.

We can now create simpler errors for the other things that can be wrong in our application.
Add the following errors in `errors.go`:

```go
var (
	ErrorEventNotValid      = ErrEventNotValid{}                    // event is not valid
	ErrorDeviceNotSupported = errors.New("device not supported")    // device not supported
	ErrorDeviceNotFound     = errors.New("device target not found") // device target not found
	ErrorEventFailed        = errors.New("failed to handle event")  // failed to handle event
)
```

### Usage of the custom errors in the services

Modify the `Server` and the `Controller` methods to return the correct errors.
In `server.go` make the code complient with the new Validate method and use the `NewErrEventNotValid` constructor.
In `controller.go` replace the return statements with an error in the Handle method with the correct custom error.

At this step, the code returns specific errors depending of what happended in the services but we don't use them in the server adapters (cli and http).

### Usage of the custom errors in the server adapters

#### http

We can now modify the `handleErrorHttp` function in `internal/adapter/driving/server/http.go` to return the correct status code depending on the error.

To check which error `err` is, we can use the `Is` function of the `errors` package](<https://pkg.go.dev/errors#Is>):

```go
// this returns true if err has an ErrorEventNotValid in it error tree
errors.Is(err, domain.ErrorEventNotValid)
```

Add cases in `handleErrorHttp` to return the correct error code based on the error.

#### cli

For the cli server adapter, the code is already implemented in `handleError` in `internal/adapter/driving/server/cli.go`.

You can check the code and test it manually:

```bash
make stack-up

go build ./cmd/cli/main.go   

# success
./main light color -n mock -r 0 -b 100 -g 200

# event not valid
./main light color -n mock -r 0 -b 100 -g 400

# target not found
./main light color -n mockkk -r 0 -b 100 -g 200
```

## Next

To see the solutions:

```bash
git checkout 6-errors-end
```

To go to the next step:

```bash
git checkout 7-end
```
