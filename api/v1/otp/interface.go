package otp

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Send(ctx context.Context, req SendReq) (SendResp, error)
	Check(ctx context.Context, req CheckReq) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
