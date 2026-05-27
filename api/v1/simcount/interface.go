package simcount

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber string, daysCount int) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
