package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestOTPSend(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterOTP)
	defer srv.Close()

	svc := otp.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Send(context.Background(), otp.SendReq{
		To: "380677770200",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.Status != "SUCCESS" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}
}

func TestOTPCheck(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterOTP)
	defer srv.Close()

	svc := otp.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), otp.CheckReq{
		SubscriberID:   "380677770200",
		ValidationCode: "4545",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.Status != "VALID" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}
}
