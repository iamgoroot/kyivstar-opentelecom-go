package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

func resolveErr(resp *http.Response) error {
	var kotError models.KotError

	err := json.NewDecoder(resp.Body).Decode(&kotError)
	if err != nil {
		return err
	}

	kotError.HTTPStatus = resp.StatusCode
	kotError.Info = parseHeaders(resp)

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
