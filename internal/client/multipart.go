package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

func PostMultipart[Resp any](ctx context.Context, r Client, path, fieldName, fileName string, file io.Reader) (Resp, models.ReqInfo, error) {
	path = fmt.Sprintf("%s/rest/%s", r.BaseURL, path)

	var buf bytes.Buffer

	mw := multipart.NewWriter(&buf)

	fw, err := mw.CreateFormFile(fieldName, fileName)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	err = mw.Close()
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, &buf)
	if err != nil {
		var resp Resp
		return resp, models.ReqInfo{}, err
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())

	return do[Resp](r, req)
}
