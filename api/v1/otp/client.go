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
	resp, info, err := client.Post[SendReq, SendResp](ctx, s.client, endpointContextPath+"/sms", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) Check(ctx context.Context, req CheckReq) (CheckResp, error) {
	resp, info, err := client.Post[CheckReq, CheckResp](ctx, s.client, endpointContextPath+"/sms/check", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}
