package simcount

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

func (s service) Check(ctx context.Context, phoneNumber string, daysCount int) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "sim-count")
	q := url.Values{"daysCount": {fmt.Sprintf("%d", daysCount)}}
	return client.Get[CheckResp](ctx, s.client, endpointPath, q)
}
