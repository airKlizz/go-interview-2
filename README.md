# Goome

## 7 - End

### Control the light over internet

The two `main.go` files have slightly been modified to use env variables instead of hard coded values.
The local values are now stored in the `.env.local` file.

> **🛠️ Action Required:**
> Run `go get github.com/joho/godotenv` to install the dependency.

> **Note:** This is not an advanced way for handling configuration in go, a commonly used tool is [viper](https://github.com/spf13/viper) but we won't use it today.

> **🛠️ Action Required:**
> Create an `.env` file with the content of the `.env.local` file to use the local stack as before.

> **🛠️ Action Required:**
> Change now the values if the environment variables with the values of the MQTT server of the real light (ask for these values) to control it via the CLI:
>
> ```bash
> go build ./cmd/cli/main.go  
> ./main light color -n mock -b 255
> ```

The real light should turn blue.
