package otp

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/verification"

type service struct {
	client client.Client
}

func (s service) Send(ctx context.Context, req SendReq) (SendResp, error) {
	return client.Post[SendReq, SendResp](ctx, s.client, endpointContextPath+"/sms", nil, req)
}

func (s service) Check(ctx context.Context, req CheckReq) (CheckResp, error) {
	return client.Post[CheckReq, CheckResp](ctx, s.client, endpointContextPath+"/sms/check", nil, req)
}
