package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/testing/local/handlers"
)

func TestMultichannelSend(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterMultichannel)
	defer srv.Close()

	svc := multichannel.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Send(context.Background(), multichannel.SendReq{
		To: "380670000200",
		SmsContent: &multichannel.SmsContent{
			Priority: 0,
			From:     "Kyivstar",
			Text:     "sms text",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.MultiMsgID == "" {
		t.Error("expected multiMsgId")
	}
}

func TestMultichannelCheck(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterMultichannel)
	defer srv.Close()

	svc := multichannel.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
	resp, err := svc.Check(context.Background(), "6badca00-2e05-42df-b7f1-4a5642e38af8")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
	if len(resp.Reports) != 2 {
		t.Errorf("expected 2 reports, got %d", len(resp.Reports))
	}
}
