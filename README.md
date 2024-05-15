# Goome

## 5 - Data Validation

### Test the Data Validation

First, let's see what happens with the current code when we try to change the light color to an invalid value.

> **🛠️ Action Required:**
> Add a test case to the `TestServer_LightChangeColor` function that tests changing the color to an invalid value (e.g., `Blue: 500`). You should see that it works, but we want the server to return an error indicating the color is not valid.

### Data Validation with [validator](https://github.com/go-playground/validator)

#### Introduction to the Package

The [validator](https://github.com/go-playground/validator) package allows performing struct validation in Golang using tags.

Here is a basic example:

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

func main() {
	validate := validator.New(validator.WithRequiredStructEnabled())
	user := User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      25,
		Username: "johndoe123",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("user not valid")
	}
	fmt.Println("user valid")
}

#### Installation

The package is already in the `go.mod` file, so run:

```bash
go mod download
```

#### Validate Event

Add validation tags to the structs in `internal/core/domain/colors.go` and `internal/core/domain/event.go`.

Validation rules:

- 0 <= Red, Green, Blue, White <= 255
- 0 <= Gain, Brightness <= 100
- 3000 <= Temp <= 6500
- Device should be "light"
- Action should be one of "on", "off", "change_color", "change_white"

Implement the following method in `event.go`:

```go
func (e *Event) Validate() error {
	panic("not implemented")
}

Use the validator package as shown in the example. Also, check the correct args for each action.

> **🛠️ Action Required:**
> When the implementation is done, run:
>
> ```bash
> go test mynewgoproject/internal/core/domain
> ```

### Usage

#### Validate Event in Server

> **🛠️ Action Required:**
> Modify the test case in `TestServer_LightChangeColor` with an invalid Blue value, set `wantErr` to true, and remove the mock's expected call. You should see an error when running the test.

> **🛠️ Action Required:**
> Modify the code in `internal/core/service/server.go` to call the validate method of the Event struct and pass the test.

> **🛠️ Action Required:**
> Run the server test to be sure it works:
>
> ```bash
> go test  mynewgoproject/internal/core/service 
> ```

#### Manual Test

To ensure it works, do a quick manual test.

```bash
# Start local stack if not already up
make stack-up

# Build the CLI
go build ./cmd/cli/main.go

# Use the CLI with invalid data
./main light color -n bedroom -r 0 -b 400 -g 200
```

It should say:

```bash
cli.go:94: failed to change color of the light
```

which is expected but not very helpful. We will next explore error management in Golang

## Next

To see the solutions:

```bash
git checkout 5-data-validation-end
```

To go to the next step:

```bash
git checkout 6-errors
```
