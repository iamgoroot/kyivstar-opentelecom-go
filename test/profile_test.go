package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestProfileGet(t *testing.T) {
	svc := profile.NewService(setupTestClient(t, handlers.RegisterProfile))

	resp, err := svc.Get(context.Background(), `{
 profile(msisdn:"380672000200"){
age
gender
 }
}`)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Data == nil || resp.Data.Profile == nil {
		t.Fatal("expected profile data")
	}

	if resp.Data.Profile.Gender != "MALE" {
		t.Errorf("unexpected gender: %s", resp.Data.Profile.Gender)
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}
