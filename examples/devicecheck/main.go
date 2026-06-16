package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
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

	svc := devicecheck.NewService(ksClient)

	resp, err := svc.Check(ctx, "380670000200", "123456789012345")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Device: %v", resp)
}
