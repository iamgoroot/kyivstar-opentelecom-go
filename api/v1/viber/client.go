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
	resp, info, err := client.Post[TransactionReq, SendResp](ctx, s.client, endpointContextPath+"/transaction", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) SendPromotionText(ctx context.Context, req PromotionTextReq) (SendResp, error) {
	resp, info, err := client.Post[PromotionTextReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) SendPromotionImage(ctx context.Context, req PromotionImageReq) (SendResp, error) {
	resp, info, err := client.Post[PromotionImageReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) SendPromotionAction(ctx context.Context, req PromotionActionReq) (SendResp, error) {
	resp, info, err := client.Post[PromotionActionReq, SendResp](ctx, s.client, endpointContextPath+"/promotion", nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) Check(ctx context.Context, msgID string) (CheckResp, error) {
	endpointPath := path.Join(endpointContextPath+"/status", msgID)
	resp, info, err := client.Get[CheckResp](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}
