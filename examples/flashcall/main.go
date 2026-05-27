package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"
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

	svc := flashcall.NewService(ksClient)

	createResp, err := svc.Create(ctx, flashcall.CreateReq{
		To: "380670000200",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Flash call created: reqID=%s", createResp.ReqID)

	checkResp, err := svc.Check(ctx, flashcall.CheckReq{
		SubscriberID:   "380670000200",
		ValidationCode: "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Flash call verified: status=%s", checkResp.Cid)
}
