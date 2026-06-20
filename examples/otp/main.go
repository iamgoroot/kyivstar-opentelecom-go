package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"
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

	svc := otp.NewService(ksClient)

	sendResp, err := svc.Send(ctx, otp.SendReq{
		To: "380677770200",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OTP sent: reqID=%s", sendResp.ReqID)

	checkResp, err := svc.Check(ctx, otp.CheckReq{
		SubscriberID:   "380677770200",
		ValidationCode: "4545",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OTP verified: status=%s", checkResp.Resource.Status)
}
