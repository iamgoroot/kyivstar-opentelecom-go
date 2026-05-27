package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcount"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestSimCount(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSimCount)
	defer srv.Close()

	svc := simcount.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "380670170200", 30)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.SimCount != 3 {
		t.Errorf("expected simCount 3, got %d", resp.Resource.SimCount)
	}
}
