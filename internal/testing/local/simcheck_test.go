package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestSimCheck(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSimCheck)
	defer srv.Close()

	svc := simcheck.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "380670010101", 24)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Resource == nil || resp.Resource.SimChanged == nil {
		t.Error("expected simChanged in resource")
	}
}
