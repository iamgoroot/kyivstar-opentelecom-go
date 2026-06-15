package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"
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

	svc := multichannel.NewService(ksClient)

	resp, err := svc.Send(ctx, multichannel.SendReq{
		To: "380670000200",
		SmsContent: &multichannel.SmsContent{
			From: "messagedesk",
			Text: "Hello via multichannel!",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent: multiMsgID=%s", resp.MultiMsgID)

	check, err := svc.Check(ctx, resp.MultiMsgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s", check.Status)
}
