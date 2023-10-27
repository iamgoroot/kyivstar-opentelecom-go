package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type requester struct {
	*http.Client
	Url     string
	Version string
}

func get[Resp any](c *requester, url string) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s/%s", c.Url, c.Version, url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		var resp Resp
		return resp, err
	}
	return do[Resp](c, req)
}

func post[Req, Resp any](c *requester, url string, body Req) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s/%s", c.Url, c.Version, url)
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		var resp Resp
		return resp, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		var resp Resp
		return resp, err
	}
	return do[Resp](c, req)
}

func do[Resp any](c *requester, req *http.Request) (Resp, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "github.com/iamgoroot/kyivstar-opentelecom-go")
	var resp Resp
	rawResp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	if rawResp.StatusCode != http.StatusOK {
		return resp, resolveErr(rawResp)
	}
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	return resp, err
}

func resolveErr(resp *http.Response) error {
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var kotError Err
	err = json.Unmarshal(all, &kotError)
	if err != nil {
		return err
	}
	kotError.HttpStatus = resp.StatusCode
	return kotError
}
