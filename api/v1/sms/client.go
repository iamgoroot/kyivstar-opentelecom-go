package sms

import (
	"context"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const (
	endpointContextPath = "v1/sms"
	batchSendEndpoint   = endpointContextPath + "/batch"
	batchCheckEndpoint  = endpointContextPath + "/status/batch"
)

type service struct {
	client client.Client
}

// Send Відправка SMS
func (s service) Send(ctx context.Context, req SendReq) (SendResp, error) {
	resp, info, err := client.Post[SendReq, SendResp](ctx, s.client, endpointContextPath, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

// SendBatch Відправка SMS (batch)
func (s service) SendBatch(ctx context.Context, req BatchSendReq) (BatchSendResp, error) {
	resp, info, err := client.Post[BatchSendReq, BatchSendResp](ctx, s.client, batchSendEndpoint, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

// Check Перевірка статусу SMS
func (s service) Check(ctx context.Context, msgID string) (CheckResp, error) {
	endpointPath := path.Join(endpointContextPath, msgID)
	resp, info, err := client.Get[CheckResp](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}

// CheckBatch Статус доставки (batch)
func (s service) CheckBatch(ctx context.Context, req BatchStatusReq) (BatchStatusResp, error) {
	resp, info, err := client.Post[BatchStatusReq, BatchStatusResp](ctx, s.client, batchCheckEndpoint, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}
