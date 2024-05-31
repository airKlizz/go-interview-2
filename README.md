# Goome

## 4 - Testing

### Overview

In Golang, testing is made with the official testing [package](https://pkg.go.dev/testing). I invite you to read the beginning of the page to have a quick overview of how to test a Go application. There are specificities but in two words:

- A function starting with `Test`, having `*testing.T` in the argument, and in a file ending with `_test.go` is a test (`func TestXxx(*testing.T)`).
- You can test in the same package as the application (without exporting the package) or in a separate `_test` package for "black box" testing.
- There are naming conventions for the names of a test function.
- Tests can be run using `go test`.

Here is a valid test sample:

```go
package abs

import "testing"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

In addition to the official `testing` package, the [`testify` package](https://github.com/stretchr/testify) is widely used for assertions and testing:

```go
package abs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        assert.Equal(t, 1, got)
    }
}
```

### First Test

Let's create the first test for our project. We can add tests for the `Controller` struct (`internal/core/service/controller.go`).

#### Mocking

The `Controller` struct is using the `Light` driven port. For testing only the code of the `Controller` struct, we need to mock the `Light` interface. There are multiple ways of creating mocks in Golang. Among them, the [`mockery` package](https://github.com/vektra/mockery) allows us to automatically generate a mock from an interface. Install it by running ([or other way](https://vektra.github.io/mockery/latest/installation/)):

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
			err := c.Handle(context.TODO(), tt.args.event)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
```

> **Note:** Table-driven tests are usually used in Golang but this is not a requirement.

To run the test with coverage, do:

```bash
go test mynewgoproject/internal/core/service -cover
```

You should see 0% coverage which is expected as there is no test case yet.

> **🛠️ Action Required:**
> You can add some test cases to increase the coverage of the service package.

For example, this is a valid test case:

```go
"OK switch on light": {
	fields: fields{
		lights: map[string]func() driven.Light{
			"bedroom": func() driven.Light {
				m := driven.NewMockLight(t)
				m.EXPECT().SwitchOn(mock.Anything).Return(nil)
				return m
			},
		},
	},
	args: args{
		event: &domain.Event{Target: "bedroom", Device: domain.Light, Action: domain.On},
	},
},
```

#### Yet Another Test

> **🛠️ Action Required:**
> Following the same principle, you can create tests for the server. Copy-paste the `controller_test.go` into `server_test.go` and adapt it to the `Server` struct.

## Next

To see the solutions:

```bash
git checkout 4-testing-end
```

To go to the next step:

```bash
git checkout 5-data-validation
```
