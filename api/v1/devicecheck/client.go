package devicecheck

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type service struct {
	client client.Client
}

func (s service) Check(ctx context.Context, phoneNumber string) (CheckRespWithResource, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "device-check")
	return client.Get[CheckRespWithResource](ctx, s.client, endpointPath, nil)
}

func (s service) CheckWithImei(ctx context.Context, phoneNumber, imei string, daysPeriod int) (CheckRespWithResource, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "device-check")
	q := url.Values{
		"imei":       {imei},
		"daysPeriod": {strconv.Itoa(daysPeriod)},
	}

	return client.Get[CheckRespWithResource](ctx, s.client, endpointPath, q)
}
