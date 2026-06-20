package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
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

	svc := profile.NewService(ksClient)

	resp, err := svc.Get(ctx, `{
 profile(msisdn:"380672000200"){
age
gender
 }
}`)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Profile: %+v", resp)
}
