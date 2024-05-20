package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		registry, err := scaleway.NewRegistryNamespace(ctx, "main", &scaleway.RegistryNamespaceArgs{
			Description: pulumi.String("Container registry for mynewgoproject"),
			IsPublic:    pulumi.Bool(false),
			Name:        pulumi.String("mynewgoproject"),
		})
		if err != nil {
			return err
		}

		ctx.Export("registryEndpoint", registry.Endpoint)
		ctx.Export("registryId", registry.ID())

		return nil
	})
}
