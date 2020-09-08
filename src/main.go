package main

import (
	"context"
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"golang.org/x/build/kubernetes/gke"
)

type gkeConfig struct {
	Project string `env:"GCP_PROJECT,required"`
	Cluster string `env:"CLUSTER_NAME,required"`
	Zone    string `env:"ZONE,required"`
	Scope   string `env:"SCOPE,required"`
}

func main() {

	// Parse server configuration
	config := gkeConfig{}
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("Env Invalid: %v", err)
	}

	ctx := context.Background()

	client, err := gke.NewClient(
		ctx,
		config.Cluster,
		gke.OptProject(config.Project),
		gke.OptZone(config.Zone))

	if err != nil {
		log.Fatalf("Could not connect to cluster: %v", err)
	}

	pods, err := client.GetPods(ctx)
	if err != nil {
		log.Fatalf("Could not query pods: %v", err)
	}
	fmt.Printf("Total pods: %d", len(pods))
}
