package profile

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Get(ctx context.Context, query string) (Resp, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
