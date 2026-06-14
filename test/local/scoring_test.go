package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/scoring"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestScoring(t *testing.T) {
	svc := scoring.NewService(setupTestClient(t, handlers.RegisterScoring))

	resp, err := svc.Check(context.Background(), "380670010103", 5)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.ScoreBal < 0.000_1 {
		t.Error("expected scoreBal above minimal threshold")
	}
}
