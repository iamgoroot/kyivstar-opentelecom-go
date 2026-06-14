package testinglocal

import (
	"context"
	"strings"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
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

func TestSMSSendError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	_, err := svc.Send(context.Background(), sms.SendReq{
		From: "messagedesk",
		To:   "invalid",
		Text: "Hello!",
	})
	if err == nil {
		t.Fatal("expected error")
	}

	var kotErr models.KotError
	if !asKotError(err, &kotErr) {
		t.Fatalf("expected KotError, got %T", err)
	}

	if kotErr.ErrorCode != 40001 {
		t.Errorf("unexpected errorCode: %d", kotErr.ErrorCode)
	}

	if !strings.Contains(kotErr.Error(), "Invalid phone number format") {
		t.Errorf("expected error message in Error(): %s", kotErr.Error())
	}

	if kotErr.HttpStatus != 400 {
		t.Errorf("unexpected httpStatus: %d", kotErr.HttpStatus)
	}
}

func TestSMSCheckError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	_, err := svc.Check(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}

	var kotErr models.KotError
	if !asKotError(err, &kotErr) {
		t.Fatalf("expected KotError, got %T", err)
	}

	if kotErr.ErrorCode != 40401 {
		t.Errorf("unexpected errorCode: %d", kotErr.ErrorCode)
	}

	if kotErr.ReqId != "err-req-id" {
		t.Errorf("unexpected reqId: %s", kotErr.ReqId)
	}

	if kotErr.HttpStatus != 404 {
		t.Errorf("unexpected httpStatus: %d", kotErr.HttpStatus)
	}
}

func TestSMSSendBatchError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	_, err := svc.SendBatch(context.Background(), sms.BatchSendReq{
		Data: map[string]sms.SendReq{"key": {From: "a", To: "b", Text: "c"}},
	})
	if err == nil {
		t.Fatal("expected error")
	}

	var kotErr models.KotError
	if !asKotError(err, &kotErr) {
		t.Fatalf("expected KotError, got %T", err)
	}

	if kotErr.ErrorCode != 40101 {
		t.Errorf("unexpected errorCode: %d", kotErr.ErrorCode)
	}
}

func TestSMSCheckBatchError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})

	_, err := svc.CheckBatch(context.Background(), sms.BatchStatusReq{
		Data: []string{"unknown"},
	})
	if err == nil {
		t.Fatal("expected error")
	}

	var kotErr models.KotError
	if !asKotError(err, &kotErr) {
		t.Fatalf("expected KotError, got %T", err)
	}

	if kotErr.ErrorCode != 50001 {
		t.Errorf("unexpected errorCode: %d", kotErr.ErrorCode)
	}
}
