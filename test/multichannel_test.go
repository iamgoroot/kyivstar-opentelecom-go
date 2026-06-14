package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestMultichannelSend(t *testing.T) {
	svc := multichannel.NewService(setupTestClient(t, handlers.RegisterMultichannel))
	var resp multichannel.SendResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Send(context.Background(), multichannel.SendReq{
			To: "380670000200",
			SmsContent: &multichannel.SmsContent{
				Priority:      0,
				From:          "Kyivstar",
				Text:          "sms text",
				MessageTtlSec: 300,
			},
			ViberContent: &multichannel.ViberContent{
				Priority:      1,
				From:          "Kyivstar",
				Text:          "viber text",
				MessageTtlSec: 300,
			},
		})
		return err
	})

	if resp.MultiMsgID == "" {
		t.Error("expected multiMsgId")
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestMultichannelCheck(t *testing.T) {
	svc := multichannel.NewService(setupTestClient(t, handlers.RegisterMultichannel))
	var resp multichannel.CheckResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Check(context.Background(), "23e39c9a-9c8d-4090-95d5-000000001200")
		return err
	})

	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}

	if len(resp.Reports) != 2 {
		t.Errorf("expected 2 reports, got %d", len(resp.Reports))
	}
}
