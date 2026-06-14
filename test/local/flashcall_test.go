package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestFlashCallCreate(t *testing.T) {
	svc := flashcall.NewService(setupTestClient(t, handlers.RegisterFlashCall))

	resp, err := svc.Create(context.Background(), flashcall.CreateReq{
		To: "380677770200",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.Status != "SUCCESS" {
		t.Errorf("unexpected status: %s", resp.Resource.Status)
	}
}

func TestFlashCallCheck(t *testing.T) {
	svc := flashcall.NewService(setupTestClient(t, handlers.RegisterFlashCall))

	resp, err := svc.Check(context.Background(), flashcall.CheckReq{
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
