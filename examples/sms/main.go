package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	ctx := context.Background()
	ksClient, err := ksOpen.NewOauthClient(ctx, ksOpen.Config{
		ServerURL:    ksOpen.Gateway,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := sms.NewService(ksClient)

	resp, err := svc.Send(ctx, sms.SendReq{
		From: "messagedesk",
		To:   "380670000200",
		Text: "Hello World!",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent: reqID=%s msgID=%s", resp.ReqID, resp.MsgID)

	check, err := svc.Check(ctx, resp.MsgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s", check.Status)
}
