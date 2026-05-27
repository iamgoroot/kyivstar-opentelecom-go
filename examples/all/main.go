package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	conf := ksOpen.Config{
		ServerUrl:    ksOpen.Gateway,
		ServerMode:   ksOpen.ServerModeMock,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	}

	ctx := context.Background()

	ksClient, err := ksOpen.NewV1Client(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	const destinationPhoneNumber = "380670000200"

	sendMsgResp, err := ksClient.SMS.Send(
		ctx,
		sms.SendReq{
			From: "messagedesk",
			To:   destinationPhoneNumber,
			Text: "Hello World!",
		},
	)
	log.Println("Sent", sendMsgResp, err)

	check, err := ksClient.SMS.Check(ctx, sendMsgResp.MsgID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Check", check.Status, err)
}
