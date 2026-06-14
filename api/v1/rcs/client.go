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
	return client.Post[RcsTextReq, SendResp](ctx, s.client, endpointContextPath+"/text", nil, req)
}

func (s service) SendSuggestion(ctx context.Context, req RcsSuggestionReq) (SendResp, error) {
	return client.Post[RcsSuggestionReq, SendResp](ctx, s.client, endpointContextPath+"/suggestion", nil, req)
}

func (s service) SendRichCard(ctx context.Context, req RcsRichCardReq) (SendResp, error) {
	return client.Post[RcsRichCardReq, SendResp](ctx, s.client, endpointContextPath+"/richcard", nil, req)
}

func (s service) Check(ctx context.Context, msgID string) (CheckResp, error) {
	endpointPath := path.Join(endpointContextPath, msgID)
	return client.Get[CheckResp](ctx, s.client, endpointPath, nil)
}
