# Goome

## 2 - Interface

We have seen how to change the color of the bulb, but this is very specific to the bulb we are using. We can use an interface to make it less specific.

### Folder Structure

In Go, there are multiple ways of organizing a project, with no specific rules except one: all the packages defined in the `internal/` folder cannot be imported by other projects. However, the following [standard](https://github.com/golang-standards/project-layout) is usually respected. In summary:

* `/cmd` contains the main applications of the project.
* `/pkg` contains public packages.
* `/internal` contains private packages.

To follow this standard, we can move our `main.go` into a new `cmd` folder:

```bash
mkdir cmd && mv main.go cmd/main.go
```

> **🛠️ Action Required:**
> Move `main.go` to the new `cmd` folder using the command above.

### Understanding Interfaces in Go

In Go, an interface is a type that specifies a method set, providing a way to define the behavior that types must implement without specifying how these methods should be implemented. An interface is a named collection of method signatures.

For example:

```go
type MyInterface interface {
    Method1(arg1 Type1) ReturnType1
    Method2(arg2 Type2) ReturnType2
}
```

A type implements an interface by implementing its methods. There is no explicit declaration of intent to implement an interface: it is satisfied implicitly.

### Example

```go
type Lamp struct {
    // fields
}

func (l Lamp) Method1(arg1 Type1) ReturnType1 {
    // method implementation
}

func (l Lamp) Method2(arg2 Type2) ReturnType2 {
    // method implementation
}

// Now, Lamp implements MyInterface because it has implemented all the methods of MyInterface
```

This implicit implementation allows for more flexible and decoupled designs, as any type that implements the necessary methods can be used wherever the interface is expected.

### Light Interface

For our light interface, we can define it as follows (`/internal/core/port/light.go`):

```go
type Light interface {
	SwitchOn(ctx context.Context) error
	SwitchOff(ctx context.Context) error
	ChangeColor(ctx context.Context, color *domain.Color) error
	ChangeWhite(ctx context.Context, white *domain.White) error
}
```

We can switch on/off a light and change the color or the white temperature.

The `Color` and `White` objects are part of our domain (`/internal/core/domain/colors.go`):

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

Now that we have the interface and the domain defined, we can implement the light interface for our Shelly MQTT. To implement an interface, the first step is to create a struct that has the methods defined in the interface:

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

A good practice is to also create a constructor that returns the interface, ensuring the struct correctly implements the interface:

```go
func NewShellyMqtt() port.Light {
	return nil
}
```

> **🛠️ Action Required:**
> Based on the previously created `main.go` file, you can complete the constructor and the methods. We want the Shelly MQTT struct to produce MQTT messages to perform the actions. The documentation of the Shelly bulb can help: [source](https://shelly-api-docs.shelly.cloud/gen1/#shelly-bulb-rgbw-mqtt).

Once the `shelly.go` file is completed, we can use the `ShellyMqtt` to change the color of the bulb to green. Replace the content of `cmd/main.go` with:

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

> **🛠️ Action Required:**
> Replace the content of `cmd/main.go` with the above code and adapt it based on your implementation of the `ShellyMqtt`. Run it with:

```bash
go run cmd/main.go
```

### Next

If your `ShellyMqtt` is working correctly, you can proceed to the next step:

```bash
git checkout 3-hexagonal
```

If you want to see the `light.go` completed:

```bash
git checkout 2-interface-end
```
