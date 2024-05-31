# Goome

## 6 - Errors

### Objective

We have seen that when an error occurs, we have no detailed information about it when using the CLI or the HTTP server. We would like to differentiate between different errors to be able to respond with the correct HTTP status code, for example.

### Errors in Golang

In Golang, creating a custom error is as simple as:

```go
var ErrNoObjectsFound = errors.New("no objects found")
```

This approach works in many cases and is widely used. However, while it allows differentiation between errors, it does not allow adding context to the error. For example, we can create a custom error for an invalid event:

```go
var ErrEventNotValid = errors.New("event not valid")
```

but we cannot add reasons why the event is not valid. To do so, we can create a new struct that implements the `error` interface:

```go
type error interface {
	Error() string
}
```

### ErrEventNotValid

Let's create this custom error.

> **🛠️ Action Required:**
> In the `internal/core/domain` folder, create a new file called `errors.go` with the following content:

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
```

We create a struct `ErrEventNotValid` with reasons being a list of strings from the event validation. The struct implements the `error` interface with the `Error() string` method. We also implement the `Is(target error) bool` method to make the `Is` function from the `errors` package work (see [doc](https://pkg.go.dev/errors#Is) for details).

You can review the `Validate` method of the `Event` struct, which has been slightly modified to return the reasons when an event is not valid.

### More Custom Errors

We have seen how to create a custom error with context stored in the error.

Now we can create simpler errors for other things that can go wrong in our application.

> **🛠️ Action Required:**
> Add the following errors in `errors.go`:

```go
var (
	ErrorEventNotValid      = ErrEventNotValid{}                    // event is not valid
	ErrorDeviceNotSupported = errors.New("device not supported")    // device not supported
	ErrorDeviceNotFound     = errors.New("device target not found") // device target not found
	ErrorEventFailed        = errors.New("failed to handle event")  // failed to handle event
)
```

### Usage of Custom Errors in the Services

> **🛠️ Action Required:**
> Modify the `Server` and `Controller` methods to return the correct errors. In `server.go`, make the code compliant with the new `Validate` method and use the `NewErrEventNotValid` constructor. In `controller.go`, replace the return statements with an error in the `Handle` method with the correct custom error.

At this step, the code returns specific errors depending on what happened in the services, but we don't use them in the server adapters (CLI and HTTP).

### Usage of Custom Errors in the Server Adapters

#### HTTP

We can now modify the `handleErrorHttp` function in `internal/adapter/driving/server/http.go` to return the correct status code depending on the error.

To check which error `err` is, we can use the `Is` function of the `errors` package:

```go
// this returns true if err has an ErrorEventNotValid in its error tree
errors.Is(err, domain.ErrorEventNotValid)
```

> **🛠️ Action Required:**
> Add cases in `handleErrorHttp` to return the correct error code based on the error.

#### CLI

For the CLI server adapter, the code is already implemented in `handleError` in `internal/adapter/driving/server/cli.go`.

> **🛠️ Action Required:**
> You can check the code and test it manually:
>
> ```bash
> make stack-up
> 
> go build ./cmd/cli/main.go   
> 
> # success
> ./main light color -n mock -r 0 -b 100 -g 200
> 
> # event not valid
> ./main light color -n mock -r 0 -b 100 -g 400
> 
> # target not found
> ./main light color -n mockkk -r 0 -b 100 -g 200
> ```

## Next

To see the solutions:

```bash
git checkout 6-errors-end
```

To go to the next step:

```bash
git checkout 7-end
```
