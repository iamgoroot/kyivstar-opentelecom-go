package sms

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Send(ctx context.Context, req SendReq) (SendResp, error)
	SendBatch(ctx context.Context, req BatchSendReq) (BatchSendResp, error)
	Check(ctx context.Context, msgID string) (CheckResp, error)
	CheckBatch(ctx context.Context, req BatchStatusReq) (BatchStatusResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
