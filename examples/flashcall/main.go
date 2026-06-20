package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"
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

	svc := flashcall.NewService(ksClient)

	createResp, err := svc.Create(ctx, flashcall.CreateReq{
		To: "380677770200",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Flash call created: reqID=%s", createResp.ReqID)

	checkResp, err := svc.Check(ctx, flashcall.CheckReq{
		SubscriberID:   "380677770200",
		ValidationCode: "4545",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Flash call verified: status=%s", checkResp.Resource.Status)
}
