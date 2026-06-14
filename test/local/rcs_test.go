package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestRCSSendText(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	resp, err := svc.SendText(context.Background(), rcs.RcsTextReq{
		From:               "messagedesk",
		To:                 "380670000200",
		ContentExtendedRcs: rcs.ContentText{Text: "Hello"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}
}

func TestRCSSendSuggestion(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	resp, err := svc.SendSuggestion(context.Background(), rcs.RcsSuggestionReq{
		From: "messagedesk",
		To:   "380670000200",
		ContentExtendedRcs: rcs.ContentSuggestion{
			Text:        "Hello",
			Suggestions: []rcs.Suggestion{{Type: "action", Text: "Go", OpenUrlAction: "https://example.com"}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}
}

func TestRCSSendRichCard(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	resp, err := svc.SendRichCard(context.Background(), rcs.RcsRichCardReq{
		From: "messagedesk",
		To:   "380670000200",
		ContentExtendedRcs: rcs.ContentRichCard{
			StandaloneCard: &rcs.StandaloneCard{
				ThumbnailImageAlignment: new("left"),
				CardOrientation:         new("horizontal"),
				CardContent: &rcs.CardContent{
					Title: "Test",
					Media: &rcs.Media{
						ThumbnailUrl: "https://example.com/thumb.png",
						FileUrl:      "https://example.com/file.png",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.MsgID == "" {
		t.Error("expected msgID")
	}
}

func TestRCSCheck(t *testing.T) {
	svc := rcs.NewService(setupTestClient(t, handlers.RegisterRCS))

	resp, err := svc.Check(context.Background(), "20200000-0000-0000-0000-380670000200")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}
