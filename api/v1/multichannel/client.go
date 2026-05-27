package multichannel

import (
	"context"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/messaging"

type service struct {
	client client.Client
}

func (s service) Send(ctx context.Context, req SendReq) (SendResp, error) {
	return client.Post[SendReq, SendResp](ctx, s.client, endpointContextPath+"/multichannel", nil, req)
}

func (s service) Check(ctx context.Context, multiMsgID string) (CheckResp, error) {
	endpointPath := path.Join(endpointContextPath+"/multichannel", multiMsgID)
	return client.Get[CheckResp](ctx, s.client, endpointPath, nil)
}
