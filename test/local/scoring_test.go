package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/scoring"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestScoring(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterScoring)
	defer srv.Close()

	svc := scoring.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	resp, err := svc.Check(context.Background(), "380670010103", 5)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Resource == nil || resp.Resource.ScoreBal != 0.08457 {
		t.Errorf("unexpected scoreBal: %f", resp.Resource.ScoreBal)
	}
}
