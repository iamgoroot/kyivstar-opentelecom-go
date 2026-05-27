package rcs

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	SendText(ctx context.Context, req RcsTextReq) (SendResp, error)
	SendSuggestion(ctx context.Context, req RcsSuggestionReq) (SendResp, error)
	SendRichCard(ctx context.Context, req RcsRichCardReq) (SendResp, error)
	Check(ctx context.Context, msgID string) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
