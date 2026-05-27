package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/lifetime"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestLifetime(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterLifetime)
	defer srv.Close()

	svc := lifetime.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "380670170200")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.LifetimeDuration == nil {
		t.Error("expected lifetimeDuration")
	} else if resp.Resource.LifetimeDuration.TimeUnit != "MONTHS" {
		t.Errorf("unexpected timeUnit: %s", resp.Resource.LifetimeDuration.TimeUnit)
	}
}
