package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/viber"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestViberSendTransaction(t *testing.T) {
	svc := viber.NewService(setupTestClient(t, handlers.RegisterViber))
	var resp viber.SendResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.SendTransaction(context.Background(), viber.TransactionReq{
			From: "messagedesk",
			To:   "380670000202",
			Text: "Hello!",
		})
		return err
	})

	if resp.Mid == "" {
		t.Error("expected mid")
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestViberSendPromotionText(t *testing.T) {
	svc := viber.NewService(setupTestClient(t, handlers.RegisterViber))
	var resp viber.SendResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.SendPromotionText(context.Background(), viber.PromotionTextReq{
			From: "messagedesk",
			To:   "380670000202",
			Text: "Hello!",
		})
		return err
	})

	if resp.Mid == "" {
		t.Error("expected mid")
	}
}

func TestViberSendPromotionImage(t *testing.T) {
	svc := viber.NewService(setupTestClient(t, handlers.RegisterViber))
	var resp viber.SendResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.SendPromotionImage(context.Background(), viber.PromotionImageReq{
			From:            "messagedesk",
			To:              "380672000202",
			ContentExtended: viber.ContentExtendedImg{Img: "https://example.com/img.png"},
		})
		return err
	})

	if resp.Mid == "" {
		t.Error("expected mid")
	}
}

func TestViberSendPromotionAction(t *testing.T) {
	svc := viber.NewService(setupTestClient(t, handlers.RegisterViber))
	var resp viber.SendResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.SendPromotionAction(context.Background(), viber.PromotionActionReq{
			From:            "messagedesk",
			To:              "380672000202",
			Text:            "Hello!",
			ContentExtended: viber.ContentExtendedAction{Img: "https://example.com/img.png", Caption: "Click", Action: "https://example.com"},
		})
		return err
	})

	if resp.Mid == "" {
		t.Error("expected mid")
	}
}

func TestViberCheck(t *testing.T) {
	svc := viber.NewService(setupTestClient(t, handlers.RegisterViber))
	var resp viber.CheckResp

	retryOnRateLimit(t, func() error {
		var err error
		resp, err = svc.Check(context.Background(), "20200000-0000-0000-0000-380670000200")
		return err
	})

	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}
