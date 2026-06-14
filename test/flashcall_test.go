package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestFlashCallCreate(t *testing.T) {
	svc := flashcall.NewService(setupTestClient(t, handlers.RegisterFlashCall))
	var resp flashcall.CreateResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Create(context.Background(), flashcall.CreateReq{
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

func TestFlashCallCheck(t *testing.T) {
	svc := flashcall.NewService(setupTestClient(t, handlers.RegisterFlashCall))
	var resp flashcall.CheckResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Check(context.Background(), flashcall.CheckReq{
			SubscriberID:   "380677770200",
			ValidationCode: "4545",
		})
		return err
	})

	if resp.Resource == nil || resp.Resource.Status != "VALID" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}
}
