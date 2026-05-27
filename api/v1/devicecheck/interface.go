package devicecheck

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	Check(ctx context.Context, phoneNumber string) (CheckRespWithResource, error)
	CheckWithImei(ctx context.Context, phoneNumber, imei string, daysPeriod int) (CheckRespWithResource, error)
}

func NewService(client client.Client) Service {
	return &service{client: client}
}
