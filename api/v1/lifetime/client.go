package lifetime

import (
	"context"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type service struct {
	client client.Client
}

func (s service) Check(ctx context.Context, phoneNumber string) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "lifetime")
	resp, info, err := client.Get[CheckResp](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}
