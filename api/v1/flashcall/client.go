package flashcall

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/verification"

type service struct {
	client client.Client
}

func (s service) Create(ctx context.Context, req CreateReq) (CreateResp, error) {
	resp, info, err := client.Post[CreateReq, CreateResp](ctx, s.client, endpointContextPath+"/flash-call", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) Check(ctx context.Context, req CheckReq) (CheckResp, error) {
	resp, info, err := client.Post[CheckReq, CheckResp](ctx, s.client, endpointContextPath+"/flash-call/check", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}
