package promo

import (
	"context"
	"io"
	"net/url"
	"path"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/promo"

type service struct {
	client client.Client
}

func (s service) CreateSMS(ctx context.Context, req CreateSMSReq) (Promo, error) {
	resp, info, err := client.Post[CreateSMSReq, Promo](ctx, s.client, endpointContextPath, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) CreateViber(ctx context.Context, req CreateViberReq) (Promo, error) {
	resp, info, err := client.Post[CreateViberReq, Promo](ctx, s.client, endpointContextPath, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) CreateRCS(ctx context.Context, req CreateRCSReq) (Promo, error) {
	resp, info, err := client.Post[CreateRCSReq, Promo](ctx, s.client, endpointContextPath, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) List(ctx context.Context, q url.Values) (ListResp, error) {
	resp, info, err := client.Get[ListResp](ctx, s.client, endpointContextPath, q)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) Get(ctx context.Context, promoUUID string) (Promo, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID)
	resp, info, err := client.Get[Promo](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) AddAudience(ctx context.Context, promoUUID string, req AddAudienceReq) (AddAudienceResp, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "audience")
	resp, info, err := client.Post[AddAudienceReq, AddAudienceResp](ctx, s.client, endpointPath, nil, req)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) AddImage(ctx context.Context, promoUUID string, fileName string, file io.Reader) (AddImageResp, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "image")
	resp, info, err := client.PostMultipart[AddImageResp](ctx, s.client, endpointPath, "file", fileName, file)
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) ChangeStatus(ctx context.Context, promoUUID, status string) (Promo, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "status", status)
	resp, info, err := client.Put[struct{}, Promo](ctx, s.client, endpointPath, nil, struct{}{})
	resp.ReqInfoGetter = info

	return resp, err
}

func (s service) GetStatistics(ctx context.Context, promoUUID string) (PromoStat, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "statistics")
	resp, info, err := client.Get[PromoStat](ctx, s.client, endpointPath, nil)
	resp.ReqInfoGetter = info

	return resp, err
}
