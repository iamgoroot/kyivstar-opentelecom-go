package testinglocal

import (
	"errors"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

func asKotError(err error, out *models.KotError) bool {
	if ke, ok := errors.AsType[models.KotError](err); ok {
		*out = ke

		return true
	}

	return false
}
