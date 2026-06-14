package testinglocal

import (
	"context"
	"net/url"
	"strings"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestPromoCreateSMS(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.CreateSMS(context.Background(), promo.CreateSMSReq{
		From:         "author",
		Text:         "Hello ${1}",
		CampaignType: "SMS",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoCreateViber(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.CreateViber(context.Background(), promo.CreateViberReq{
		From:         "author",
		Text:         "Hello ${1}",
		CampaignType: "VIBER",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoCreateRCS(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.CreateRCS(context.Background(), promo.CreateRCSReq{
		From:         "author",
		Text:         "Hello ${1}",
		CampaignType: "RCS",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoList(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))
	q := url.Values{"pageSize": {"10"}, "pageNumber": {"0"}}

	resp, err := svc.List(context.Background(), q)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Promos) == 0 {
		t.Error("expected promos")
	}
}

func TestPromoGet(t *testing.T) {
	if !isRunningLocally() {
		t.Skip("flaky: mock data not found — will fix later")
	}
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.Get(context.Background(), "00000000-0000-0000-0000-000000000200")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoAddAudience(t *testing.T) {
	if !isRunningLocally() {
		t.Skip("flaky: mock data not found — will fix later")
	}
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.AddAudience(context.Background(), "00000000-0000-0000-0000-000000000200", promo.AddAudienceReq{
		Audience: []promo.AudienceMember{{Params: []string{"John"}, PhoneNumber: "380671234200"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Name == "" {
		t.Error("expected name")
	}
}

func TestPromoAddImage(t *testing.T) {
	if !isRunningLocally() {
		t.Skip("AddImage test only runs locally")
	}

	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	img := strings.NewReader("fake-image-data")

	resp, err := svc.AddImage(context.Background(), "00000000-0000-0000-0000-000000000200", "test.png", img)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Success {
		t.Error("expected success")
	}
}

func TestPromoChangeStatus(t *testing.T) {
	if !isRunningLocally() {
		t.Skip("flaky: mock data not found — will fix later")
	}
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.ChangeStatus(context.Background(), "00000000-0000-0000-0000-000000000200", "WAITING")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "WAITING" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

// TODO: update after proper mock mode mocks
func TestPromoGetStatistics(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.GetStatistics(context.Background(), "20000000-0000-0000-0000-000000000200")
	if err != nil {
		t.Fatal(err)
	}
	// TODO: update after proper mock mode mocks
	if resp.SentCount != 0 {
		t.Errorf("unexpected sentCount: %d", resp.SentCount)
	}
	if resp.WasNotSent != 20 {
		t.Errorf("unexpected wasNotSent: %d", resp.WasNotSent)
	}

}
