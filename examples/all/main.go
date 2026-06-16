package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	var conf ksOpen.Config
	if err := conf.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ksClient, err := ksOpen.NewV1Client(ctx, &conf)
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
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sent", sendMsgResp)

	check, err := ksClient.SMS.Check(ctx, sendMsgResp.MsgID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Checking status. Status:", check.Status, err)
}
