package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcheck"
)

func main() {
	ctx := context.Background()
	var conf ksOpen.Config
	if err := conf.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	ksClient, err := ksOpen.NewOauthClient(ctx, &conf)
	if err != nil {
		log.Fatal(err)
	}

	svc := simcheck.NewService(ksClient)

	resp, err := svc.Check(ctx, "380670010101", 24)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SIM check result: %v", resp)
}
