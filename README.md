# Goome

## 5 - Data validation

### Test the data validation

In Golang, testing is made with the official testing [package](https://pkg.go.dev/testing). I invite you to read the beginning of the page to have a quick overview of how to test a Go application. There are specificities but in two words:

- A function starting with `Test`, having `*testing.T` in the argument, and in a file ending with `_test.go` is a test (`func TestXxx(*testing.T)`).
- You can test in the same package as the application (without exporting the package) or in a separate `_test` package for "black box" testing.
- There are naming conventions for the names of a test function.
- Tests can be run using `go test`.

### Data validation with [validator](https://github.com/go-playground/validator)

#### Introduction to the package

The [validator](https://github.com/go-playground/validator) package allows to perform struct validation in Golang using tags in the structs.

Here is a basic example of use of the package:

```go
package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Define a sample struct to validate
type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"min=0,max=130"`
	Username string `validate:"required,min=3,max=16"`
}

#White: {
	Temp:       uint & >=3000 & <=6500
	Brightness: uint & >=0 & <=100
}
```

It defines constraints for the Go structs contain in the `colors.go` file.

```go
package abs

```cue
package cue

Event: {
	Target: string
	Device: "light"
}

### First Test

Let's create the first test for our project. We can add tests for the `Controller` struct (`internal/core/service/controller.go`).

#### Mocking

The `Controller` struct is using the `Light` driven port. For testing only the code of the `Controller` struct, we need to mock the `Light` interface. There are multiple ways of creating mocks in Golang. Among them, the [`mockery` package](https://github.com/vektra/mockery) allows us to automatically generate a mock from an interface. Install it by running ([or other way](https://vektra.github.io/mockery/latest/installation/))

```bash
go install github.com/vektra/mockery/v2@v2.43.0
```

and generate the mocks using:

```bash
mockery
```

> **🛠️ Action Required:**
> Install the package and generate the mocks. You can have a look at the `.mockery.yaml` file to see the configuration of the tool.

#### Table-Driven Test

Now that we have the mock we need, we can create the test for the `Controller`.

> **🛠️ Action Required:**
> Create the `controller_test.go` file in the same folder with the following content:

```go
package service

import (
	"context"
	"errors"
	"testing"

	"mynewgoproject/internal/core/domain"
	"mynewgoproject/internal/core/port/driven"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_Handle(t *testing.T) {
	type fields struct {
		lights map[string]func() driven.Light
	}
	type args struct {
		event *domain.Event
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr bool
	}{
		// include tests here
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewController()
			for name, light := range tt.fields.lights {
				c.WithLight(name, light())
			}
		}
	} |
	{
		Action: "change_white"
		Args: {
			ChangeWhiteArgs: {
				White: #White
			}
		}
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("user not valid")
	}
	fmt.Println("user valid")
}
```

> **Note:** Table-driven tests are usually used in Golang but this is not a requirement.

- `Device` can only be `"light"` as we support only this device.
- Each `Action` should come with its `Args`

#### Validation

CUE provides a Go [package](https://pkg.go.dev/cuelang.org/go@v0.8.2/encoding/gocode) to generate Go validation code to the struct of a package based on a CUE schema.

We will take advantage of `go generate` to generate validation code for the domain structs based on the CUE schema automaticaly.

The code for this is not very interesting, therefore it is already present in the project:

- `internal/core/domain/cue.mod` file to define the CUE module name
- `internal/core/domain/doc.go` file containing the call to special comment for generation. The comment in the file simply says: run `go run ../../utils/gen/cue.go` when `go generate` is ran on this package.
- `internal/utils/gen/cue.go` file that contains the code to use the CUE packages to generate the validation methods for the structs.

With these files in the project, by running

```bash
go mod download
```

You should see 0% coverage which is expected as there is no test case yet.

> **🛠️ Action Required:**
> You can add some test cases to increase the coverage of the service package.

For example, this is a valid test case

- 0 <= Red, Green, Blue, White <= 255
- 0 <= Gain, Brightness <= 100
- 3000 <= Temp <= 6500
- Device should be one if "light"
- Action should be one if "on off change_color change_white"

Then, add the following method to the `event.go` file and implement it:

```go
func (e *Event) Validate() error {
	panic("not implemented")
}
```

Use the validator package like in the example but check also the correct args are present for the action (i.e. if the action is change_color then the change color args should be present)

When the implementation is done, run:

```bash
go test mynewgoproject/internal/core/domain
```

to make sure the implementation is correct.

### Usage

#### Validate Event in server

Let's start by modifying the test case we added just before in `TestServer_LightChangeColor` with an invalid Blue value to make it match what we want.
We want the `LightChangeColor` method to return an error when calling with an invalid color.
Therefore change the `wantErr` bool of the test case to true and remove the expected call to the mock.

When running the test you should see something similar to:

```bash
$ go test mynewgoproject/internal/core/service       
--- FAIL: TestServer_LightChangeColor (0.00s)
    --- FAIL: TestServer_LightChangeColor/KO_color_not_valid (0.00s)
        server_test.go:72: 
                Error Trace:    goome/internal/core/service/server_test.go:72
                Error:          An error is expected but got nil.
                Test:           TestServer_LightChangeColor/KO_color_not_valid
    mock_Light.go:220: PASS:    ChangeColor(string,*domain.Color)
    mock_Light.go:220: PASS:    ChangeColor(string,*domain.Color)
FAIL
FAIL    mynewgoproject/internal/core/service    0.339s
FAIL
```

Modify the code in `internal/core/service/server.go` to call the validate method of the Event struct and pass the test.

#### Manual test

To make sure it works, we can do a quick manual test.

```bash
# start local stack if not already up
make stack-up

# build the CLI
go build ./cmd/cli/main.go

# use the CLI with invalid data
./main light color -n bedroom -r 0 -b 400 -g 200
```

#### Yet Another Test

> **🛠️ Action Required:**
> Following the same principle, you can create tests for the server. Copy-paste the `controller_test.go` into `server_test.go` and adapt it to the `Server` struct.

## Next

To see the solutions:

```bash
git checkout 5-data-validation-end
```

To go to the next step:

```bash
git checkout 6-errors
```
