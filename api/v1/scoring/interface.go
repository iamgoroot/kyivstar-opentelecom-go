package scoring

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber string, scoreFormula int) (CheckResp, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
