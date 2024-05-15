# Goome

## 5 - Data validation

### Test the data validation

In Golang, testing is made with the official testing [package](https://pkg.go.dev/testing). I invite you to read the beginning of the page to have a quick overview of how to test a Go application. There are specificities but in two words:

- A function starting with `Test`, having `*testing.T` in the argument, and in a file ending with `_test.go` is a test (`func TestXxx(*testing.T)`).
- You can test in the same package as the application (without exporting the package) or in a separate `_test` package for "black box" testing.
- There are naming conventions for the names of a test function.
- Tests can be run using `go test`.

### Data validation with CUE

#### Introduction to CUE

CUE is a [data validation language](https://cuelang.org/docs/introduction/) written in Go which make it relatively well integrated with Go (despite not perfect yet).

One of the usage of CUE is to define a schema in the CUE language and to validate data based on this schema.
It can validates JSON, YAML, and Go structs which is what we will use.

#### Schema

We define the schemas in the `internal/core/domain` folder.
Create a `cue` folder that will contains our CUE schemas:

```bash
mkdir internal/core/domain/cue 
```

Then create a `colors.cue` file that contains:

```cue
package cue

#Color: {
	Red:   uint & >=0 & <=255
	Green: uint & >=0 & <=255
	Blue:  uint & >=0 & <=255
	White: uint & >=0 & <=255
	Gain:  uint & >=0 & <=100
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

With these files in the project, by running:

```bash
# install new packages
go mod tidy 

# generate code
go generate ./...
```

You should see 0% coverage which is expected as there is no test case yet.

> **🛠️ Action Required:**
> You can add some test cases to increase the coverage of the service package.

For example, this is a valid test case

When you run again the test in `server_test.go` it should fail as it returns an error.
Adapt the test to expect an error when the input data is not valid.

Now our service only accepts valid input data.

### Usage

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
