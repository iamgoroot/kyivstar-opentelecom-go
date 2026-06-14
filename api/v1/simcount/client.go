package simcount

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

func (s service) Check(ctx context.Context, phoneNumber string, daysCount int) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "sim-count")
	q := url.Values{"daysCount": {strconv.Itoa(daysCount)}}

	resp, info, err := client.Get[CheckResp](ctx, s.client, endpointPath, q)
	resp.ReqInfoGetter = info

	return resp, err
}
