package ksopentelecom

import (
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type (
	SMS = sms.Service
)
type V1Client struct {
	SMS
}

func createV1Client(ksClient client.Client) (V1Client, error) {
	return V1Client{
		SMS: sms.NewService(ksClient),
	}, nil
}
