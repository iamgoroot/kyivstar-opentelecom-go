package lifetime

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber string) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
