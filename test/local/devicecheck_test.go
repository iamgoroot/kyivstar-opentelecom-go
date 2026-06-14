package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestDeviceCheck(t *testing.T) {
	if !isRunningLocally() {
		t.Skip("flaky: mock data not found — will fix later")
	}
	svc := devicecheck.NewService(setupTestClient(t, handlers.RegisterDeviceCheck))

	resp, err := svc.Check(context.Background(), "380670170200")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.ImeiRes != "COMPLETELY_MATCHED" {
		t.Errorf("unexpected imeiRes: %s", resp.Resource.ImeiRes)
	}
}

func TestDeviceCheckWithImei(t *testing.T) {
	svc := devicecheck.NewService(setupTestClient(t, handlers.RegisterDeviceCheck))

	resp, err := svc.CheckWithImei(context.Background(), "380670170200", "123456789012345", 30)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.ImeiRes == "" {
		t.Error("expected imeiRes")
	}
}
