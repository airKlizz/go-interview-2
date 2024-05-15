# Goome

## 5 - Data validation

### Test the data validation

First of all we can see what happen with the current code when we want to change the color of the light with a color not valid.

Add a test case to the `TestServer_LightChangeColor` function that tests to change the color with an invalid color (`Blue: 500` for example).
We see that it works but we would prefer the server to return some kind of error that says the color is not valid.

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

In a similar way, we can define constraints for the `Event` struct.
The schema is a bit more complex (`internal/core/domain/cue/event.cue`):

```cue
package cue

Event: {
	Target: string
	Device: "light"
}

Event:
	{
		Action: "on"
	} |
	{
		Action: "off"
	} |
	{
		Action: "change_color"
		Args: {
			ChangeColorArgs: {
				Color: #Color
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

The schema says:

* `Device` can only be `"light"` as we support only this device.
* Each `Action` should come with its `Args`

#### Validation

CUE provides a Go [package](https://pkg.go.dev/cuelang.org/go@v0.8.2/encoding/gocode) to generate Go validation code to the struct of a package based on a CUE schema.

We will take advantage of `go generate` to generate validation code for the domain structs based on the CUE schema automaticaly.

The code for this is not very interesting, therefore it is already present in the project:

* `internal/core/domain/cue.mod` file to define the CUE module name
* `internal/core/domain/doc.go` file containing the call to special comment for generation. The comment in the file simply says: run `go run ../../utils/gen/cue.go` when `go generate` is ran on this package.
* `internal/utils/gen/cue.go` file that contains the code to use the CUE packages to generate the validation methods for the structs.

With these files in the project, by running:

```bash
# install new packages
go mod tidy 

# generate code
go generate ./...
```

it will create the `internal/core/domain/cue_gen.go` file which contains the `Validate()` methods for the `Event` struct.

### Validate events

Now that the `Event` struct has a `Validate() error` method, use it in the `Server` methods to return an error if the created event is not valid.

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
