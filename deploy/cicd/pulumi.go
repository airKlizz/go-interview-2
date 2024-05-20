package cicd

import (
	"context"

	"dagger.io/dagger"
)

func NewScwPulumiContainer(ctx context.Context, client *dagger.Client, directory, scwAccessKey, scwSecretKey, scwDefaultOrganizationId, scwDefaultProjectId, pulumuAccessToken string) *dagger.Container {
	return client.Container().
		From("pulumi/pulumi-go").
		WithExec([]string{"/bin/bash", "-c", "pulumi plugin install resource scaleway --server github://api.github.com/pulumiverse"}).
		WithSecretVariable("PULUMI_ACCESS_TOKEN", client.SetSecret("PULUMI_ACCESS_TOKEN", pulumuAccessToken)).
		WithSecretVariable("SCW_ACCESS_KEY", client.SetSecret("SCW_ACCESS_KEY", scwAccessKey)).
		WithSecretVariable("SCW_SECRET_KEY", client.SetSecret("SCW_SECRET_KEY", scwSecretKey)).
		WithEnvVariable("SCW_DEFAULT_ORGANIZATION_ID", scwDefaultOrganizationId).
		WithEnvVariable("SCW_DEFAULT_PROJECT_ID", scwDefaultProjectId).
		WithDirectory("/pulumi/projects", client.Host().Directory(directory))
}
