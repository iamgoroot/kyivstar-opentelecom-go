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
	return client.Post[CreateSMSReq, Promo](ctx, s.client, endpointContextPath, nil, req)
}

func (s service) CreateViber(ctx context.Context, req CreateViberReq) (Promo, error) {
	return client.Post[CreateViberReq, Promo](ctx, s.client, endpointContextPath, nil, req)
}

func (s service) CreateRCS(ctx context.Context, req CreateRCSReq) (Promo, error) {
	return client.Post[CreateRCSReq, Promo](ctx, s.client, endpointContextPath, nil, req)
}

func (s service) List(ctx context.Context, q url.Values) (ListResp, error) {
	return client.Get[ListResp](ctx, s.client, endpointContextPath, q)
}

func (s service) Get(ctx context.Context, promoUUID string) (Promo, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID)
	return client.Get[Promo](ctx, s.client, endpointPath, nil)
}

func (s service) AddAudience(ctx context.Context, promoUUID string, req AddAudienceReq) (AddAudienceResp, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "audience")
	return client.Post[AddAudienceReq, AddAudienceResp](ctx, s.client, endpointPath, nil, req)
}

func (s service) AddImage(ctx context.Context, promoUUID string, fileName string, file io.Reader) (AddImageResp, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "image")
	return client.PostMultipart[AddImageResp](ctx, s.client, endpointPath, "file", fileName, file)
}

func (s service) ChangeStatus(ctx context.Context, promoUUID, status string) (Promo, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "status", status)
	return client.Put[struct{}, Promo](ctx, s.client, endpointPath, nil, struct{}{})
}

func (s service) GetStatistics(ctx context.Context, promoUUID string) (PromoStat, error) {
	endpointPath := path.Join(endpointContextPath, promoUUID, "statistics")
	return client.Get[PromoStat](ctx, s.client, endpointPath, nil)
}
