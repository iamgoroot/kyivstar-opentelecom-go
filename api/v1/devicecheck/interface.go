package devicecheck

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber, imei string) (CheckRespWithResource, error)
	CheckWithImei(ctx context.Context, phoneNumber string, daysPeriod int) (CheckRespWithResource, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
