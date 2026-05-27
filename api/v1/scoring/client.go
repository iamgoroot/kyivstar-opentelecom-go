package scoring

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

func (s service) Check(ctx context.Context, phoneNumber string, scoreFormula int) (CheckResp, error) {
	endpointPath := path.Join("v1/subscribers", phoneNumber, "scoring")
	q := url.Values{"scoreFormula": {strconv.Itoa(scoreFormula)}}

	return client.Get[CheckResp](ctx, s.client, endpointPath, q)
}
