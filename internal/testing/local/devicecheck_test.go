package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestDeviceCheck(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterDeviceCheck)
	defer srv.Close()

	svc := devicecheck.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "380670170200")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.ImeiRes != "COMPLETELY_MATCHED" {
		t.Errorf("unexpected imeiRes: %s", resp.Resource.ImeiRes)
	}
}

func TestDeviceCheckWithImei(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterDeviceCheck)
	defer srv.Close()

	svc := devicecheck.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.CheckWithImei(context.Background(), "380670170200", "123456789012345", 30)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.ImeiRes != "COMPLETELY_MATCHED" {
		t.Errorf("unexpected imeiRes: %s", resp.Resource.ImeiRes)
	}
}
