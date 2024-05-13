# Goome

## 3 - Hexagonal

### Vocabulary

We can continue building the application and now see how to controlle the bulb.
If we think our application following the hexagonal architecture, the light interface is a driven port and the Shelly MQTT implementation of it is an adapter.
We now need a driving port to controlle the light that we call `Server`:

```go
type Server interface {
	LightOn(ctx context.Context, name string) error
	LightOff(ctx context.Context, name string) error
	LightChangeColor(ctx context.Context, name string, color *domain.Color) error
	LightChangeWhite(ctx context.Context, name string, white *domain.White) error
}
```

We can imagine multiple adapters of the `Server` port: a http server, a CLI, a website.

To make the link between the driving and the driven ports, we have the controller service:

```go
type Controller struct {
	lights     map[string]driven.Light
}

func (c *Controller) Handle(ctx context.Context, event *domain.Event) error {}
```

The controller knows the lights available and can handle events.
A event is part of the domain definition of the application:

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

You can have a look to all the new files in `internal`.
The services are already implemented.

### Server adapters

Now the vocabulary of the application is defined, we can work on implementing adapters of the server port to actually controlle the bulb.

We will do two adapters, a HTTP server and a CLI

#### HTTP server

In go, there are plenty of HTTP server, one of the most used is [Gin](https://github.com/gin-gonic/gin).

A structure of an implementation of a Gin server is made in the `internal/adapter/driving/server/http.go` file.
There is a structure `HttpServer` which contains a `Server` and which can `Run()` to start the HTTP server.

The objective here is to get the data from the HTTP call and to call the server with this data.

🫵 You can implement the Gin handler function of the file. The function should bind the data and call the server. Here is a useful [link](https://gin-gonic.com/docs/examples/bind-query-or-post/) to the Gin documentation.

When the `HttpServer` is implemented, you can look to the `cmd/http/main.go` file which does the dependency injections and start the HTTP server, and run:

```bash
go run cmd/http/main.go
```

to start the HTTP server.

You can change the bulb color with a curl command:

```bash
curl -X POST --data '{"name": "mock", "color": {"blue": 200, "gain": 100}}' http://localhost:8080/light/color
```

The color of the bulb on [http://localhost:3333/](http://localhost:3333/) or using `make bulb-color` should be blue.

#### CLI

For building a CLI able to controlle the bulb, we can use the [Cobra](https://github.com/spf13/cobra) package.
Similar to what we did for the HTTP server, we create a structure `CliServer` which contains a `Server` and which can `Run()` to execute the CLI.

The code is verbose for writing a CLI, therefore it is already written.
You can look at it in `internal/adapter/driving/server/cli.go`.

🫵 You can copy the `cmd/http/main.go` file into `cmd/cli/main.go` file and adapt it to not run the `HttpServer` but the `CliServer`.

Run:

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
