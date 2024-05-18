# Goome

## 5 - Data validation

### Test the data validation

First of all we can see what happen with the current code when we want to change the color of the light with a color not valid.

Add a test case to the `TestServer_LightChangeColor` function that tests to change the color with an invalid color (`Blue: 500` for example).
We see that it works but we would prefer the server to return some kind of error that says the color is not valid.

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
```

#### Installation

The package has been added to the `go.mod` file, so you can simply run:

```bash
go mod download
```

#### Validate Event

First of all to use it, add the correct tags to the structs in the `internal/core/domain/colors.go` and `internal/core/domain/event.go` files.

Respect the following rules:

* 0 <= Red, Green, Blue, White <= 255
* 0 <= Gain, Brightness <= 100
* 3000 <= Temp <= 6500
* Device should be one if "light"
* Action should be one if "on off change_color change_white"

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

It should says:

```bash
cli.go:94: failed to change color of the light
```

which is ok but not really helpful.
We will see next one way to manage errors in Golang.

## Next

To see the solutions:

```bash
git checkout 5-data-validation-end
```

To go to the next step:

```bash
git checkout 6-errors
```
