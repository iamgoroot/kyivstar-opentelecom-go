package main

import (
	"context"
	"fmt"
	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
)

func main() {
	conf := ksOpen.Config{
		ServerUrl:    ksOpen.Gateway,
		ServerMode:   ksOpen.ServerModeMock,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	}
	ctx := context.Background()
	ksClient := ksOpen.New(ctx, conf)
	const destinationPhoneNumber = "380670000200"
	//Send msg
	sendMsgResp, err := ksClient.Send(
		ksOpen.SmsSendReq{
			From: "messagedesk",
			To:   destinationPhoneNumber,
			Text: "Hello World!",
		})
	fmt.Println("Sent", sendMsgResp, err)

	// Check Status
	check, err := ksClient.Check(sendMsgResp.MsgId)
	fmt.Println("Check", check.Status, err)

	// Scoring
	scoring, err := ksClient.Scoring(destinationPhoneNumber, 0)
	fmt.Println("Scored:", scoring, err)

	//Verify sim
	sim, err := ksClient.VerifySim(destinationPhoneNumber, ksOpen.VerifySimReq{
		ActivationHours: 48,
	})
	fmt.Printf("Verify sim: changed=%d, active=%d, err=%v", sim.SimChanged, sim.IsActive, err)
}
