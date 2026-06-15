package promo

import (
	"context"
	"io"
	"net/url"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
	CreateSMS(ctx context.Context, req CreateSMSReq) (Promo, error)
	CreateViber(ctx context.Context, req CreateViberReq) (Promo, error)
	CreateRCS(ctx context.Context, req CreateRCSReq) (Promo, error)
	List(ctx context.Context, q url.Values) (ListResp, error)
	Get(ctx context.Context, promoUUID string) (Promo, error)
	AddAudience(ctx context.Context, promoUUID string, req AddAudienceReq) (AddAudienceResp, error)
	AddImage(ctx context.Context, promoUUID string, fileName string, file io.Reader) (AddImageResp, error)
	ChangeStatus(ctx context.Context, promoUUID, status string) (Promo, error)
	GetStatistics(ctx context.Context, promoUUID string) (Stat, error)
}

func NewService(c client.Client) Service {
	return &service{client: c}
}
