package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestRCSSendText(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterRCS)
	defer srv.Close()

	svc := rcs.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
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
	srv := handlers.NewServer(handlers.RegisterRCS)
	defer srv.Close()

	svc := rcs.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
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
	srv := handlers.NewServer(handlers.RegisterRCS)
	defer srv.Close()

	svc := rcs.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.SendRichCard(context.Background(), rcs.RcsRichCardReq{
		From: "messagedesk",
		To:   "380670000200",
		ContentExtendedRcs: rcs.ContentRichCard{
			StandaloneCard: &rcs.StandaloneCard{
				CardOrientation: new("horizontal"),
				CardContent: &rcs.CardContent{
					Title: "Test",
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
	srv := handlers.NewServer(handlers.RegisterRCS)
	defer srv.Close()

	svc := rcs.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "test-msg-id")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}
