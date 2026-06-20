package flashcall

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Create(ctx context.Context, req CreateReq) (CreateResp, error)
	Check(ctx context.Context, req CheckReq) (CheckResp, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
