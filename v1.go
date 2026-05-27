package ksopentelecom

import "context"

func NewV1Client(ctx context.Context, conf Config) (V1Client, error) {
	ksClient, err := createOauthClient(ctx, conf)
	if err != nil {
		return V1Client{}, err
	}
	return createV1Client(ksClient)
}
