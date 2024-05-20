# Goome

## 7 - Cloud

### Control the light over internet

The two `main.go` files have slightly been modified to use env variables instead of hard coded values.
The local values are now stored in the `.env.local` file.

> This is not an advanced way for handling configuration in go, a commonly used tool is [viper](https://github.com/spf13/viper) but we won't use it today.

Create an `.env` file with the content of the `.env.local` file to use the local stack as before.

Change now the values if the environment variables with the values of the MQTT server of the real light (ask for these values) to control it via the CLI:

```bash
go build ./cmd/cli/main.go  
./main light color -n mock -b 255
```

The real light should turn blue.

### Deploy the HTTP server on Scaleway

The objective of this part is to build a container image of the HTTP server and to deploy it on Scalway using the serveless container service.

#### Build the container

We will use [dagger](https://dagger.io/) to do the build using Go.
