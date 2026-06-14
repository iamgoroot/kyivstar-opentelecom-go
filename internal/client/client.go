package client

import (
	"context"
	"net/http"
	"net/url"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

//go:generate mockgen -source=client.go -destination=mock_doer_test.go -package=client

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Client  Doer
	BaseUrl string
}

func Get[Resp any](ctx context.Context, r Client, path string, q url.Values) (Resp, models.ReqInfo, error) {
	fullURL := composeUrl(r.BaseUrl, path, q)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, http.NoBody)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	return do[Resp](r, req)
}

func Post[Req, Resp any](ctx context.Context, r Client, url string, q url.Values, body Req) (Resp, models.ReqInfo, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPost, url, q, body)
}

func Put[Req, Resp any](ctx context.Context, r Client, url string, q url.Values, body Req) (Resp, models.ReqInfo, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPut, url, q, body)
}
