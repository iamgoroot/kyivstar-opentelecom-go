package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestProfileGet(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterProfile)
	defer srv.Close()

	svc := profile.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Get(context.Background(), `{ profile(msisdn:"380672000200") { age gender } }`)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Data == nil || resp.Data.Profile == nil {
		t.Fatal("expected profile data")
	}
	if resp.Data.Profile.Gender != "MALE" {
		t.Errorf("unexpected gender: %s", resp.Data.Profile.Gender)
	}
}
