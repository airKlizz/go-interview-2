package cicd

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func BuildCli(ctx context.Context, client *dagger.Client, address string) error {
	return build(ctx, client, address, "cli")
}

func BuildHttp(ctx context.Context, client *dagger.Client, address string) error {
	return build(ctx, client, address, "http")
}

func build(ctx context.Context, client *dagger.Client, address string, app string) error {
	src := client.Host().Directory(".")
	builder := client.Container().
		From("golang:1.21").
		WithDirectory("/src", src).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-o", "app", "./cmd" + app})

	prodImage := client.Container().
		From("alpine").
		WithFile("/bin/app", builder.File("/src/app")).
		WithEntrypoint([]string{"/bin/app"})

	addr, err := prodImage.Publish(ctx, address)
	if err != nil {
		return err
	}
	fmt.Printf("Container image accessible at the following address: %s", addr)
	return nil
}
