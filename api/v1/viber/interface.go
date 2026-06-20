package viber

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	SendTransaction(ctx context.Context, req TransactionReq) (SendResp, error)
	SendPromotionText(ctx context.Context, req PromotionTextReq) (SendResp, error)
	SendPromotionImage(ctx context.Context, req PromotionImageReq) (SendResp, error)
	SendPromotionAction(ctx context.Context, req PromotionActionReq) (SendResp, error)
	Check(ctx context.Context, msgID string) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
