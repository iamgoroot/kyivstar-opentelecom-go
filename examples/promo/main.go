package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"
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

	svc := promo.NewService(ksClient)

	p, err := svc.CreateSMS(ctx, promo.CreateSMSReq{
		From:         "author",
		Text:         "Hello ${1}",
		CampaignType: "SMS",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created promo: id=%s", p.ID)
}
