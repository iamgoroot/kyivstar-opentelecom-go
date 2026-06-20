package profile

import (
	"context"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/subscribers"

type service struct {
	client client.Client
}

func (s service) Get(ctx context.Context, query string) (Resp, error) {
	resp, info, err := client.Post[QueryReq, Resp](ctx, s.client, endpointContextPath+"/profile", nil, QueryReq{Query: query})
	resp.ReqInfoGetter = info

	return resp, err
}
