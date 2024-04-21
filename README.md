# Goome

## 3 - Hexagonal

### Vocabulary

We can continue building the application and now see how to control the bulb. If we think of our application following the hexagonal architecture, the light interface is a driven port, and the Shelly MQTT implementation of it is an adapter. We now need a driving port to control the light, which we call `Server`:

```go
type Server interface {
	LightOn(ctx context.Context, name string) error
	LightOff(ctx context.Context, name string) error
	LightChangeColor(ctx context.Context, name string, color *domain.Color) error
	LightChangeWhite(ctx context.Context, name string, white *domain.White) error
}
```

We can imagine multiple adapters of the `Server` port: an HTTP server, a CLI, a website.

To make the link between the driving and the driven ports, we have the controller service:

```go
type Controller struct {
	lights     map[string]driven.Light
}

func (c *Controller) Handle(ctx context.Context, event *domain.Event) error {}
```

The controller knows the lights available and can handle events. An event is part of the domain definition of the application

```go
type Event struct {
	Target string
	Device Device
	Action Action
	Args   *Args
}

type Device string

const (
	Light  = "light"
)

type Action string

const (
	On          = "on"
	Off         = "off"
	ChangeColor = "change_color"
	ChangeWhite = "change_white"
)

type Args struct {
	OnArgs          *OnArgs
	OffArgs         *OffArgs
	ChangeColorArgs *ChangeColorArgs
	ChangeWhiteArgs *ChangeWhiteArgs
}

type OnArgs struct{}

type OffArgs struct{}

type ChangeColorArgs struct {
	Color *Color
}

type ChangeWhiteArgs struct {
	White *White
}
```

You can have a look at all the new files in `internal`. The services are already implemented so you can skip it if you want.

### Server Adapters

Now that the vocabulary of the application is defined, we can work on implementing adapters of the server port to actually control the bulb.

We will create two adapters: an HTTP server and a CLI.

#### HTTP Server

In Go, we can use the standard `net/http` package to create an HTTP server.

A structure of an implementation of an HTTP server is made in the `internal/adapter/driving/server/http.go` file. There is a structure `HttpServer` which contains a `Server` and which can `Run()` to start the HTTP server.

The objective here is to get the data from the HTTP call and to call the server with this data.

> **🛠️ Action Required:**
> Implement the handler functions in the `http.go` file. The functions should parse the incoming request data and call the appropriate server methods. Here is an example of how to [parse JSON requests](https://go.dev/play/p/y_LWUROls8j).

When the `HttpServer` is implemented, you can look at the `cmd/http/main.go` file, which handles the dependency injections and starts the HTTP server. Then, run

```bash
go run cmd/http/main.go
```

to start the HTTP server.

> **🛠️ Action Required:**
> You can change the bulb color with a curl command:

```bash
curl -X POST --data '{"name": "mock", "color": {"blue": 200, "gain": 100}}' http://localhost:8080/light/color
```

The color of the bulb on [http://localhost:3333/](http://localhost:3333/) or using `make bulb-color` should be blue.

#### CLI

To build a CLI capable of controlling the bulb, we can use the [Cobra](https://github.com/spf13/cobra) package. Similar to what we did for the HTTP server, we create a structure `CliServer` which contains a `Server` and which can `Run()` to execute the CLI.

The code is verbose for writing a CLI, so it is already written. You can look at it in `internal/adapter/driving/server/cli.go`.

> **🛠️ Action Required:**
> Copy the `cmd/http/main.go` file into `cmd/cli/main.go` and adapt it to run the `CliServer` instead of the `HttpServer`.

> **🛠️ Action Required:**
> Run:

```bash
go run cmd/cli/main.go light color -n mock -r 100 -b 200 -g 100
```

This changes the color of the bulb to purple using the CLI.

## Next

To see the solutions:

```bash
git checkout 3-hexagonal-end
```

To go to the next step:

```bash
git checkout 4-testing
```
