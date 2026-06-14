package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcheck"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestSimCheck(t *testing.T) {
	svc := simcheck.NewService(setupTestClient(t, handlers.RegisterSimCheck))

	resp, err := svc.Check(context.Background(), "380670010101", 24)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.SimChanged == nil {
		t.Error("expected simChanged in resource")
	}
}
