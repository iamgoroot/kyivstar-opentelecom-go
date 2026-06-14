package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

const (
	estimatedPayloadSize = 1024
	maxPooledBufferSize  = estimatedPayloadSize * 16
)

var bufferPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, estimatedPayloadSize))
	},
}

func do[Resp any](r Client, req *http.Request) (Resp, models.ReqInfo, error) {
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", "github.com/iamgoroot/kyivstar-opentelecom-go")

	var resp Resp

	rawResp, err := r.Client.Do(req)
	if err != nil {
		return resp, models.ReqInfo{}, err
	}

	defer func() {
		_, _ = io.Copy(io.Discard, rawResp.Body)
		_ = rawResp.Body.Close()
	}()

	info := parseHeaders(rawResp)

	if rawResp.StatusCode >= 300 {
		return resp, info, resolveErr(rawResp)
	}

	err = json.NewDecoder(rawResp.Body).Decode(&resp)

	return resp, info, err
}

func doWithBody[Req, Resp any](ctx context.Context, r Client, method string, path string, q url.Values, body Req) (Resp, models.ReqInfo, error) {
	fullURL := composeUrl(r.BaseUrl, path, q)

	bodyBuf := bufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()

	defer func() {
		if bodyBuf.Cap() < maxPooledBufferSize {
			bufferPool.Put(bodyBuf)
		}
	}()

	err := json.NewEncoder(bodyBuf).Encode(body)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyBuf)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	return do[Resp](r, req)
}

func composeUrl(base, url string, q url.Values) string {
	if len(q) > 0 {
		return fmt.Sprintf("%s/rest/%s?%s", base, url, q.Encode())
	}
	return fmt.Sprintf("%s/rest/%s", base, url)
}
