package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestDeviceCheck(t *testing.T) {
	svc := devicecheck.NewService(setupTestClient(t, handlers.RegisterDeviceCheck))
	var resp devicecheck.CheckRespWithResource

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Check(context.Background(), "380670170200")
		return err
	})

	if resp.Resource == nil || resp.Resource.ImeiRes != "COMPLETELY_MATCHED" {
		t.Errorf("unexpected imeiRes: %s", resp.Resource.ImeiRes)
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestDeviceCheckWithImei(t *testing.T) {
	svc := devicecheck.NewService(setupTestClient(t, handlers.RegisterDeviceCheck))
	var resp devicecheck.CheckRespWithResource

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.CheckWithImei(context.Background(), "380670170200", "123456789012345", 30)
		return err
	})

	if resp.Resource == nil || resp.Resource.ImeiRes == "" {
		t.Error("expected imeiRes")
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}
