package ksopentelecom

import (
	"context"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

func NewOauthClient(ctx context.Context, conf Config) (client.Client, error) {
	authConf := &clientcredentials.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		TokenURL:     conf.ServerURL + "/idp/oauth2/token",
		EndpointParams: url.Values{
			"grant_type": []string{"client_credentials"},
		},
	}
	oauthClient := oauth2.NewClient(ctx, authConf.TokenSource(ctx))

	var serverMode string
	if conf.ServerMode != "" {
		serverMode = fmt.Sprint("/", conf.ServerMode)
	}

	result, err := url.JoinPath(conf.ServerURL, serverMode)
	if err != nil {
		return client.Client{}, err
	}

	ksClient := client.Client{
		Client:  oauthClient,
		BaseURL: result,
	}

	return ksClient, nil
}
