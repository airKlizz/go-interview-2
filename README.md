# Goome

## 2 - Interface

We have seen how to change the color of the bulb but this is very specific to the bulb we are using.
We can use an interface to make is less specific.

### Folder structure

In Go, there are multiple ways of organizing a project and no specific rules expect one: all the package defined in the `internal/` folder cannot be imported by other projects.
However, the following [standard](https://github.com/golang-standards/project-layout) is usually respected.
In two words:

* `/cmd` contains the main applications of the project
* `/pkg` contains public packages
* `/internal` contains private packages

To respect the standard, we can move our `main.go` in a new `cmd` folder:

```bash
mkdir cmd && mv main.go cmd/main.go
```

### Light interface

An interface in Go is a type which is a named collection of method signatures.
For our light interface we can define it as follows (`/internal/core/port/light.go`):

```go
type Light interface {
	SwitchOn(ctx context.Context) error
	SwitchOff(ctx context.Context) error
	ChangeColor(ctx context.Context, color *domain.Color) error
	ChangeWhite(ctx context.Context, white *domain.White) error
}
```

We can switch on/off a light and change the color or the white temperature.

The `Color` and `White` objects are parts of our domain (`/internal/core/domain/colors.go`):

```go
type Color struct {
	Red   int32
	Green int32
	Blue  int32
	White int32
	Gain  int32
}

type White struct {
	Temp       int32
	Brightness int32
}
```

Now that we have the interface and the domain defined, we can implement the light interface for our Shelly MQTT.
To implement an interface, the first step is to create a struct that has the methods defined in the interface:

```go
type ShellyMqtt struct {
}

func (c *ShellyMqtt) ChangeColor(ctx context.Context, color *domain.Color) error {
	panic("not implemented")
}

func (c *ShellyMqtt) ChangeWhite(ctx context.Context, white *domain.White) error {
	panic("not implemented")
}

func (c *ShellyMqtt) SwitchOff(ctx context.Context) error {
	panic("not implemented")
}

func (c *ShellyMqtt) SwitchOn(ctx context.Context) error {
	panic("not implemented")
}
```

A good practice is to also create a constructor that returns the interface, this makes sure the struct implements well the interface:

```go
func NewShellyMqtt() port.Light {
	return nil
}
```

🫵 Based on the previously made `main.go` file, you can complete the constructor and the methods. We want the Shelly MQTT struct to produce MQTT messages to perform the actions.
The documentation of the Shelly bulb can help: [source](https://shelly-api-docs.shelly.cloud/gen1/#shelly-bulb-rgbw-mqtt).

Once the `shelly.go` file completed, we can use the `ShellyMqtt` to change the color of the bulb to green.
Replace the content of  `cmd/main.go` with:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"mynewgoproject/internal/adapter/light"
	"mynewgoproject/internal/core/domain"
)

func main() {
	bulb := light.NewShellyMqtt()
	err := bulb.ChangeColor(context.Background(), &domain.Color{
		Red:   0,
		Green: 255,
		Blue:  0,
		White: 0,
		Gain:  100,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to change color: %w", err))
	}
	log.Println("successfully changed color")
}
```

You can adapt the code based on your implementation of the `ShellyMqtt` and run it with:

```bash
go run cmd/main.go
```

## Next

If your `ShellyMqtt` is working correctly, you can directly go to the next step:

```bash
git checkout 3-hexagonal
```

If you want to see the `light.go` completed:

```bash
git checkout 2-interface-end
```
