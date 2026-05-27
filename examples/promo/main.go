package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"
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

	svc := promo.NewService(ksClient)

	p, err := svc.CreateSMS(ctx, promo.CreateSMSReq{
		From:         "messagedesk",
		Text:         "Hello from promo!",
		CampaignType: "sms",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created promo: id=%s", p.ID)
}
