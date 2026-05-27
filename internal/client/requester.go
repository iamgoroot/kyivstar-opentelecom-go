package client

import (
	"bytes"
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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		var resp Resp
		return resp, err
	}
	return do[Resp](r, req)
}

func Post[Req, Resp any](ctx context.Context, r Client, url string, q url.Values, body Req) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s", r.BaseUrl, url)
	if len(q) > 0 {
		url += "?" + q.Encode()
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var resp Resp
		return resp, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		var resp Resp
		return resp, err
	}
	return do[Resp](r, req)
}

func do[Resp any](r Client, req *http.Request) (Resp, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "github.com/iamgoroot/kyivstar-opentelecom-go")
	var resp Resp
	rawResp, err := r.Client.Do(req)
	if err != nil {
		return resp, err
	}

	defer rawResp.Body.Close()

	if rawResp.StatusCode >= 300 {
		return resp, resolveErr(rawResp)
	}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	return resp, err
}

func resolveErr(resp *http.Response) error {
	var kotError models.Err
	err := json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {
		return err
	}
	kotError.HttpStatus = resp.StatusCode
	return kotError
}
