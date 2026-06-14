package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
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

func doWithBody[Req, Resp any](ctx context.Context, r Client, method string, url string, q url.Values, body Req) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s", r.BaseUrl, url)

	if len(q) > 0 {
		url += "?" + q.Encode()
	}

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
		return resp, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyBuf)
	if err != nil {
		var resp Resp
		return resp, err
	}

	return do[Resp](r, req)
}

func do[Resp any](r Client, req *http.Request) (Resp, error) {
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", "github.com/iamgoroot/kyivstar-opentelecom-go")

	var resp Resp

	rawResp, err := r.Client.Do(req)
	if err != nil {
		return resp, err
	}

	defer func() {
		_, _ = io.Copy(io.Discard, rawResp.Body)
		_ = rawResp.Body.Close()
	}()

	if rawResp.StatusCode >= 300 {
		return resp, resolveErr(rawResp)
	}

	err = json.NewDecoder(rawResp.Body).Decode(&resp)

	return resp, err
}

func PostMultipart[Resp any](ctx context.Context, r Client, url, fieldName, fileName string, file io.Reader) (Resp, error) {
	url = fmt.Sprintf("%s/rest/%s", r.BaseUrl, url)

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	fw, err := mw.CreateFormFile(fieldName, fileName)
	if err != nil {
		var resp Resp
		return resp, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		var resp Resp
		return resp, err
	}

	err = mw.Close()
	if err != nil {
		var resp Resp
		return resp, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		var resp Resp
		return resp, err
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())

	return do[Resp](r, req)
}
