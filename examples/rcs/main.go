package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"
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

	svc := rcs.NewService(ksClient)

	resp, err := svc.SendText(ctx, rcs.TextReq{
		From: "messagedesk",
		To:   "380670000200",
		ContentExtendedRcs: rcs.ContentText{
			Text: "Hello via RCS!",
		},
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
