# Goome

## 1 - Getting Started

### Local Environment

In this Hands-on, we will create a simple home automation system capable of controlling a [smart bulb via MQTT](https://shelly-api-docs.shelly.cloud/gen1/#shelly-bulb-rgbw-mqtt).

To start the local environment, run:

```bash
make stack-up
```

This starts an MQTT broker, an observability stack, and a mock of the bulb implemented in the `test/shelly_bulb_mock` folder.

You can see the bulb color by running:

```bash
make bulb-color
```

or by accessing [http://localhost:3333/](http://localhost:3333/).

### Change the Color

Let's write our first line of Go to produce an MQTT message that will change the color of the bulb.

#### Create Go Project

First, when starting a Go project, you need to run the command:

```bash
go mod init mynewgoproject
```

> Typically, the name given to a Go project is the path of the repository. In my case, it would be `gitlab.prod.aws.wescale.fr:remi.calizzano/goome`.

Then create a file `main.go` with the following content:

```go
package main

func main() {
    // your code here
}
```

To run the code, use:

```bash
go run main.go
```

#### MQTT in Go

Install the MQTT package using:

```bash
go get -u github.com/eclipse/paho.mqtt.golang
```

Now we can update the `main.go` file to produce a simple MQTT message:

```go
package main

import (
    "fmt"
    "log"
    "time"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

    // Init the MQTT client
    opts := MQTT.NewClientOptions()
    opts.AddBroker("tcp://localhost:1883")
    opts.SetUsername("mynewgoproject")
    client := MQTT.NewClient(opts)
    if token := client.Connect(); !token.WaitTimeout(time.Second) || token.Error() != nil {
        log.Fatal(fmt.Errorf("failed to init mqtt client: %w", token.Error()))
    }

    // Publish message to the bulb topic
    data := `
    {
        "mode": "color",    
        "red": 255,           
        "green": 0,         
        "blue": 0,        
        "gain": 100,        
        "brightness": 0,  
        "white": 0,         
        "temp": 0,       
        "effect": 0,        
        "turn": "on",       
        "transition": 500  
    }
    `
    token := client.Publish("shellies/shellycolorbulb-mock/color/0/set", 0, false, data)
    if !token.WaitTimeout(time.Second) || token.Error() != nil {
        log.Fatal(fmt.Errorf("failed to publish data to mqtt: %w", token.Error()))
    }

    log.Println("message successfully sent to mqtt")
}
```

When running the Go code again, you should see a success log message. When you check the bulb color again, it should be red. You can play with the values in the data and see the bulb color change.

### Next

If everything works well, proceed to the next step:

```bash
git checkout 2-interface
```

or see what you should have at the end of this step:

```bash
git checkout 1-getting-started-end
```
