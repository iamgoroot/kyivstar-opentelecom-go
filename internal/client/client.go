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
	BaseURL string
}

func Get[Resp any](ctx context.Context, r Client, path string, q url.Values) (Resp, models.ReqInfo, error) {
	fullURL := composeURL(r.BaseURL, path, q)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, http.NoBody)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	return do[Resp](r, req)
}

func Post[Req, Resp any](ctx context.Context, r Client, path string, q url.Values, body Req) (Resp, models.ReqInfo, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPost, path, q, body)
}

func Put[Req, Resp any](ctx context.Context, r Client, path string, q url.Values, body Req) (Resp, models.ReqInfo, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPut, path, q, body)
}
