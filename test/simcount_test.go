package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcount"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestSimCount(t *testing.T) {
	svc := simcount.NewService(setupTestClient(t, handlers.RegisterSimCount))

	resp, err := svc.Check(context.Background(), "380670170200", 30)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.SimCount == 0 {
		t.Error("expected non-zero simCount")
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}
