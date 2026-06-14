package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
)

func main() {
	ctx := context.Background()
	ksClient, err := ksOpen.NewOauthClient(ctx, ksOpen.Config{
		ServerUrl:    ksOpen.Gateway,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := profile.NewService(ksClient)

	resp, err := svc.Get(ctx, `{ profile(msisdn:"380670000200") { age gender } }`)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Profile: %+v", resp)
}
