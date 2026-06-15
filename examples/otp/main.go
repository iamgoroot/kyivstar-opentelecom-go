package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"
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

	svc := otp.NewService(ksClient)

	sendResp, err := svc.Send(ctx, otp.SendReq{
		To: "380670000200",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OTP sent: reqID=%s", sendResp.ReqID)

	checkResp, err := svc.Check(ctx, otp.CheckReq{
		SubscriberID:   "380670000200",
		ValidationCode: "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OTP verified: status=%s", checkResp.Resource.Status)
}
