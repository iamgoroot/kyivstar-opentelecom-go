package testinglocal

import (
	"context"
	"strings"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func TestSMSSend(t *testing.T) {
	svc := sms.NewService(setupTestClient(t, handlers.RegisterSMS))

	var (
		resp sms.SendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.Send(context.Background(), sms.SendReq{
			From: "messagedesk",
			To:   "380670000200",
			Text: "Hello World!",
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

func TestSMSCheck(t *testing.T) {
	svc := sms.NewService(setupTestClient(t, handlers.RegisterSMS))

	var (
		resp sms.CheckResp
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

func TestSMSSendBatch(t *testing.T) {
	svc := sms.NewService(setupTestClient(t, handlers.RegisterSMS))

	var (
		resp sms.BatchSendResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.SendBatch(context.Background(), sms.BatchSendReq{
			Data: map[string]sms.SendReq{
				"uniqueMsgKey1": {From: "messagedesk", To: "380670000200", Text: "Hello World!"},
				"uniqueMsgKey2": {From: "messagedesk", To: "380670000201", Text: "Hello again!"},
			},
		})

		return err
	})

	if len(resp.Data) != 2 {
		t.Errorf("expected 2 items, got %d", len(resp.Data))
	}

	if resp.Data["uniqueMsgKey1"].MsgID == "" {
		t.Error("expected msgID for uniqueMsgKey1")
	}
}

func TestSMSCheckBatch(t *testing.T) {
	svc := sms.NewService(setupTestClient(t, handlers.RegisterSMS))

	var (
		resp sms.BatchStatusResp
		err  error
	)

	retryOnRateLimit(t, func() error {
		resp, err = svc.CheckBatch(context.Background(), sms.BatchStatusReq{
			Data: []string{"20200000-0000-0000-0000-380670000200"},
		})

		return err
	})

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

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseURL: srv.URL})

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

	if kotErr.HTTPStatus != 400 {
		t.Errorf("unexpected httpStatus: %d", kotErr.HTTPStatus)
	}
}

func TestSMSCheckError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseURL: srv.URL})

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

	if kotErr.ReqID != "err-req-id" {
		t.Errorf("unexpected reqId: %s", kotErr.ReqID)
	}

	if kotErr.HTTPStatus != 404 {
		t.Errorf("unexpected httpStatus: %d", kotErr.HTTPStatus)
	}
}

func TestSMSSendBatchError(t *testing.T) {
	srv := handlers.NewServer(handlers.RegisterSMSErrors)
	defer srv.Close()

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseURL: srv.URL})

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

	svc := sms.NewService(client.Client{Client: srv.Client(), BaseURL: srv.URL})

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
