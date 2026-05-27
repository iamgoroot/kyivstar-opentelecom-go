package simcheck

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

func (s service) Check(ctx context.Context, phoneNumber string, checkPeriod int) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "sim-check-antifraud")
	q := url.Values{"checkPeriod": {strconv.Itoa(checkPeriod)}}

	return client.Get[CheckResp](ctx, s.client, endpointPath, q)
}
