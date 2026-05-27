package viber

import (
	"context"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/viber"

type service struct {
	client client.Client
}

func (s service) SendTransaction(ctx context.Context, req TransactionReq) (SendResp, error) {
	return client.Post[TransactionReq, SendResp](ctx, s.client, endpointContextPath+"/transaction", nil, req)
}

func (s service) SendPromotionText(ctx context.Context, req PromotionTextReq) (SendResp, error) {
	return client.Post[PromotionTextReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
}

func (s service) SendPromotionImage(ctx context.Context, req PromotionImageReq) (SendResp, error) {
	return client.Post[PromotionImageReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
}

func (s service) SendPromotionAction(ctx context.Context, req PromotionActionReq) (SendResp, error) {
	return client.Post[PromotionActionReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
}

func (s service) Check(ctx context.Context, msgID string) (resp CheckResp, err error) {
	endpointPath := path.Join(endpointContextPath+"/status", msgID)
	return client.Get[CheckResp](ctx, s.client, endpointPath, nil)
}
