package simcheck

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type service struct {
	client client.Client
}

func (s service) Check(ctx context.Context, phoneNumber string, checkPeriod int) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "sim-check-antifraud")
	q := url.Values{"checkPeriod": {fmt.Sprintf("%d", checkPeriod)}}
	return client.Get[CheckResp](ctx, s.client, endpointPath, q)
}
