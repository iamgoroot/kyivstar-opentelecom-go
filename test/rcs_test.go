package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestRCSSendText(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	var (
		resp rcs.SendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.SendText(context.Background(), rcs.TextReq{
			From:               "messagedesk",
			To:                 "380670000200",
			ContentExtendedRcs: rcs.ContentText{Text: "Hello"},
		})

		return err
	})

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}

	info := resp.GetReqInfo()
	if info.RequestID == "" {
		t.Error("expected RequestID")
	}
}

func TestRCSSendSuggestion(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	var (
		resp rcs.SendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.SendSuggestion(context.Background(), rcs.SuggestionReq{
			From: "messagedesk",
			To:   "380670000200",
			ContentExtendedRcs: rcs.ContentSuggestion{
				Text:        "Hello",
				Suggestions: []rcs.Suggestion{{Type: "action", Text: "Go", OpenURLAction: "https://example.com"}},
			},
		})

		return err
	})

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}
}

func TestRCSSendRichCard(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	var (
		resp rcs.SendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.SendRichCard(context.Background(), rcs.RichCardReq{
			From: "messagedesk",
			To:   "380670000200",
			ContentExtendedRcs: rcs.ContentRichCard{
				StandaloneCard: &rcs.StandaloneCard{
					ThumbnailImageAlignment: new("left"),
					CardOrientation:         new("horizontal"),
					CardContent: &rcs.CardContent{
						Title: "Test",
						Media: &rcs.Media{
							ThumbnailURL: "https://example.com/thumb.png",
							FileURL:      "https://example.com/file.png",
						},
					},
				},
			},
		})

		return err
	})

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}
}

func TestRCSCheck(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	var (
		resp rcs.CheckResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.Check(context.Background(), "20200000-0000-0000-0000-380670000200")

		return err
	})

	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}
