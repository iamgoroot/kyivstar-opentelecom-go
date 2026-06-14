package rcs

import (
	"context"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/rcs"

type service struct {
	client client.Client
}

func (s service) SendText(ctx context.Context, req RcsTextReq) (SendResp, error) {
	resp, info, err := client.Post[RcsTextReq, SendResp](ctx, s.client, endpointContextPath+"/text", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) SendSuggestion(ctx context.Context, req RcsSuggestionReq) (SendResp, error) {
	resp, info, err := client.Post[RcsSuggestionReq, SendResp](ctx, s.client, endpointContextPath+"/suggestion", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) SendRichCard(ctx context.Context, req RcsRichCardReq) (SendResp, error) {
	resp, info, err := client.Post[RcsRichCardReq, SendResp](ctx, s.client, endpointContextPath+"/richcard", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) Check(ctx context.Context, msgID string) (CheckResp, error) {
	endpointPath := path.Join(endpointContextPath, msgID)
	resp, info, err := client.Get[CheckResp](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}
