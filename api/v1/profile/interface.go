package profile

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Get(ctx context.Context, query string) (ProfileResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
