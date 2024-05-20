package main

import (
	"context"
	"fmt"
	"log"
	"mynewgoproject/deploy/cicd"
	"os"

	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect to dagger: %s", err.Error()))
	}
	defer client.Close()

	endpoint, id, err := cicd.CreateRegistry(
		ctx,
		client,
		os.Getenv("SCW_ACCESS_KEY"),
		os.Getenv("SCW_SECRET_KEY"),
		os.Getenv("SCW_DEFAULT_ORGANIZATION_ID"),
		os.Getenv("SCW_DEFAULT_PROJECT_ID"),
		os.Getenv("PULUMI_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to created registry: %s", err.Error()))
	}
	log.Println(endpoint, id)
}
