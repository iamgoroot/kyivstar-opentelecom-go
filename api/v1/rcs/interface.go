package rcs

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	SendText(ctx context.Context, req TextReq) (SendResp, error)
	SendSuggestion(ctx context.Context, req SuggestionReq) (SendResp, error)
	SendRichCard(ctx context.Context, req RichCardReq) (SendResp, error)
	Check(ctx context.Context, msgID string) (CheckResp, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
