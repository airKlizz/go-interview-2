package cicd

import (
	"context"

	"dagger.io/dagger"
)

func CreateRegistry(ctx context.Context, client *dagger.Client, scwAccessKey, scwSecretKey, scwDefaultOrganizationId, scwDefaultProjectId, pulumuAccessToken string) (string, string, error) {
	pulumiContainer := NewScwPulumiContainer(ctx, client, "deploy/infra/registry", scwAccessKey, scwSecretKey, scwDefaultOrganizationId, scwDefaultProjectId, pulumuAccessToken)
	registryOutput := pulumiContainer.WithExec([]string{"/bin/bash", "-c", "pulumi stack select dev -c && pulumi up -y --skip-preview"})
	endpoint, err := registryOutput.WithExec([]string{"/bin/bash", "-c", "pulumi stack output registryEndpoint -s dev"}).Stdout(ctx)
	if err != nil {
		return "", "", err
	}
	id, err := registryOutput.WithExec([]string{"/bin/bash", "-c", "pulumi stack output registryId -s dev"}).Stdout(ctx)
	if err != nil {
		return "", "", err
	}
	return endpoint, id, nil
}
