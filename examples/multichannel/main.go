package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"
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

	svc := multichannel.NewService(ksClient)

	resp, err := svc.Send(ctx, multichannel.SendReq{
		To: "380670000200",
		SmsContent: &multichannel.SmsContent{
			Priority:      0,
			From:          "Kyivstar",
			Text:          "sms text message",
			MessageTTLSec: 30,
		},
		ViberContent: &multichannel.ViberContent{
			Priority:      1,
			From:          "Kyivstar",
			PromoType:     "PROMOTIONAL",
			Text:          "viber text message",
			MessageTTLSec: 30,
			Img:           new("http://img.png"),
			Caption:       new("caption text"),
			Action:        new("http://google.com"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent: multiMsgID=%s", resp.MultiMsgID)

	check, err := svc.Check(ctx, "23e39c9a-9c8d-4090-95d5-000000001200")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s, Reports: %d", check.Status, len(check.Reports))
}
