package testinglocal

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"net/url"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestPromoCreateSMS(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	var (
		resp promo.Promo
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.CreateSMS(context.Background(), promo.CreateSMSReq{
			From:         "author",
			Text:         "Hello ${1}",
			CampaignType: "SMS",
		})

		return err
	})

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestPromoCreateViber(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	var (
		resp promo.Promo
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.CreateViber(context.Background(), promo.CreateViberReq{
			From:         "author",
			Text:         "Hello ${1}",
			CampaignType: "VIBER",
		})

		return err
	})

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoCreateRCS(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	var (
		resp promo.Promo
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.CreateRCS(context.Background(), promo.CreateRCSReq{
			From:         "author",
			Text:         "Hello ${1}",
			CampaignType: "RCS",
		})

		return err
	})

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoList(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))
	q := url.Values{"pageSize": {"10"}, "pageNumber": {"0"}}

	var (
		resp promo.ListResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.List(context.Background(), q)

		return err
	})

	if len(resp.Promos) == 0 {
		t.Error("expected promos")
	}
}

func TestPromoGet(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.Get(context.Background(), "10000000-0000-0000-0000-000000000200")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "DRAFT" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoAddAudience(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.AddAudience(context.Background(), "10000000-0000-0000-0000-000000000200", promo.AddAudienceReq{
		Audience: []promo.AudienceMember{{Params: []string{"John", "Smith"}, PhoneNumber: "380671234200"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Name == "" {
		t.Error("expected name")
	}
}

func TestPromoAddImage(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	var buf bytes.Buffer

	err := png.Encode(&buf, img)
	if err != nil {
		t.Error(err)
	}

	resp, err := svc.AddImage(context.Background(), "10000000-0000-0000-0000-000000000200", "test.png", &buf)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Success {
		t.Error("expected success")
	}
}

func TestPromoChangeStatus(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	resp, err := svc.ChangeStatus(context.Background(), "20000000-0000-0000-0000-000000000200", "WAITING")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "WAITING" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestPromoGetStatistics(t *testing.T) {
	svc := promo.NewService(setupTestClient(t, handlers.RegisterPromo))

	var (
		resp promo.Stat
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.GetStatistics(context.Background(), "20000000-0000-0000-0000-000000000200")

		return err
	})

	if resp.SentCount != 0 {
		t.Errorf("unexpected sentCount: %d", resp.SentCount)
	}

	if resp.WasNotSent != 20 {
		t.Errorf("unexpected wasNotSent: %d", resp.WasNotSent)
	}
}
