package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/viber"
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
