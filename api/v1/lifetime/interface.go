package lifetime

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber string) (CheckResp, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
