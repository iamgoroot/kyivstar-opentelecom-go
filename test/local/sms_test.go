package testinglocal

import (
	"context"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestSMSSend(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMS)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	resp, err := svc.Send(context.Background(), sms.SendReq{
		From: "messagedesk",
		To:   "380670000200",
		Text: "Hello World!",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.MsgID != "20200000-0000-0000-0000-380670000200" {
		t.Errorf("unexpected msgID: %s", resp.MsgID)
	}
}

func TestSMSCheck(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMS)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	resp, err := svc.Check(context.Background(), "20200000-0000-0000-0000-380670000200")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "delivered" {
		t.Errorf("unexpected status: %s", resp.Status)
	}
}

func TestSMSSendBatch(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMS)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	resp, err := svc.SendBatch(context.Background(), sms.BatchSendReq{
		Data: map[string]sms.SendReq{
			"uniqueMsgKey1": {From: "messagedesk", To: "380670000200", Text: "Hello World!"},
			"uniqueMsgKey2": {From: "messagedesk", To: "380670000201", Text: "Hello again!"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("expected 2 items, got %d", len(resp.Data))
	}
	if resp.Data["uniqueMsgKey1"].MsgID != "20200000-0000-0000-0000-380670000200" {
		t.Errorf("unexpected msgID: %s", resp.Data["uniqueMsgKey1"].MsgID)
	}
}

func TestSMSCheckBatch(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMS)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	resp, err := svc.CheckBatch(context.Background(), sms.BatchStatusReq{
		Data: []string{"20200000-0000-0000-0000-380670000200"},
	})
	if err != nil {
		t.Fatal(err)
	}

	item, ok := resp.Data["20200000-0000-0000-0000-380670000200"]
	if !ok {
		t.Fatal("expected key in data map")
	}
	if item.Status != "delivered" {
		t.Errorf("unexpected status: %s", item.Status)
	}
}
