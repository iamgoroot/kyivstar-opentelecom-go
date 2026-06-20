package client

import (
	"net/http"
	"strconv"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

func parseIntHeader(h http.Header, key string) int {
	v := h.Get(key)
	if v == "" {
		return 0
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return n
}

func parseHeaders(rawResp *http.Response) models.ReqInfo {
	h := rawResp.Header

	return models.ReqInfo{
		RateLimit: models.RateLimit{
			Limit:     parseIntHeader(h, "X-Rate-Limit-Limit"),
			PeriodSec: parseIntHeader(h, "X-Rate-Limit-Period-Sec"),
			Remaining: parseIntHeader(h, "X-Rate-Limit-Remaining"),
			Reset:     parseIntHeader(h, "X-Rate-Limit-Reset"),
		},
		Tariffication: models.Tariffication{
			Units: parseIntHeader(h, "X-Reserved-Tarification-Units"),
		},
		RequestID: h.Get("X-Request-Id"),
	}
}
