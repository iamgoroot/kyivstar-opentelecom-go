package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestOTPSend(t *testing.T) {
	svc := otp.NewService(setupTestClient(t, handlers.RegisterOTP))

	var (
		resp otp.SendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.Send(context.Background(), otp.SendReq{
			To: "380677770200",
		})

		return err
	})

	if resp.Resource == nil || resp.Resource.Status != "SUCCESS" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestOTPCheck(t *testing.T) {
	svc := otp.NewService(setupTestClient(t, handlers.RegisterOTP))

	var (
		resp otp.CheckResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.Check(context.Background(), otp.CheckReq{
			SubscriberID:   "380677770200",
			ValidationCode: "4545",
		})

		return err
	})

	if resp.Resource == nil || resp.Resource.Status != "VALID" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}
}
