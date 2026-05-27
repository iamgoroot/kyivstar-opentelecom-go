package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/viber"
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

	svc := viber.NewService(ksClient)

	resp, err := svc.SendTransaction(ctx, viber.TransactionReq{
		From: "messagedesk",
		To:   "380670000200",
		Text: "Hello via Viber!",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent: reqID=%s mid=%s", resp.ReqID, resp.Mid)

	check, err := svc.Check(ctx, resp.Mid)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s", check.Status)
}
