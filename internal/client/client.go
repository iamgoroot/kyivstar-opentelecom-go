package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

type Client struct {
	Client  *http.Client
	BaseUrl string
}

func Get[Resp any](ctx context.Context, r Client, url string, q url.Values) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s", r.BaseUrl, url)
	if len(q) > 0 {
		url += "?" + q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		var resp Resp
		return resp, err
	}

	return do[Resp](r, req)
}

func Post[Req, Resp any](ctx context.Context, r Client, url string, q url.Values, body Req) (Resp, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPost, url, q, body)
}

func Put[Req, Resp any](ctx context.Context, r Client, url string, q url.Values, body Req) (Resp, error) {
	return doWithBody[Req, Resp](ctx, r, http.MethodPut, url, q, body)
}

func resolveErr(resp *http.Response) error {
	var kotError models.KotError

	err := json.NewDecoder(resp.Body).Decode(&kotError)
	if err != nil {
		return err
	}

	kotError.HttpStatus = resp.StatusCode

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return fmt.Errorf("%w: %w", models.ErrBadRequestParams, kotError)
	case http.StatusUnauthorized:
		return fmt.Errorf("%w: %w", models.ErrUnauthorized, kotError)
	case http.StatusForbidden:
		return fmt.Errorf("%w: %w", models.ErrForbidden, kotError)
	case http.StatusNotFound:
		return fmt.Errorf("%w: %w", models.ErrNotFound, kotError)
	case http.StatusRequestEntityTooLarge:
		return fmt.Errorf("%w: %w", models.ErrPayloadTooLarge, kotError)
	case http.StatusUnprocessableEntity:
		return fmt.Errorf("%w: %w", models.ErrUnprocessable, kotError)
	case http.StatusTooManyRequests:
		return fmt.Errorf("%w: %w", models.ErrRateLimitExceeded, kotError)
	default:
		return fmt.Errorf("%w: %w", models.ErrInternalServer, kotError)
	}
}
