package client

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
)

const version = "v1beta"

type serverMode string

const (
	Gateway                      = "https://api-gateway.kyivstar.ua"
	ServerModeLive    serverMode = ""
	ServerModeMock    serverMode = "mock"
	ServerModeSandbox serverMode = "sandbox"
)

type Config struct {
	ServerUrl    string     `json:"serverUrl,omitempty" yaml:"serverUrl"`
	ClientID     string     `json:"clientId,omitempty" yaml:"clientId"`
	ClientSecret string     `json:"clientSecret,omitempty" yaml:"clientSecret"`
	ServerMode   serverMode `json:"serverMode,omitempty" yaml:"serverMode"`
	DebugEnabled bool       `json:"debugEnabled" yaml:"debugEnabled"`
}

func New(ctx context.Context, conf Config) Client {
	authConf := &clientcredentials.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		TokenURL:     fmt.Sprintf("%s/idp/oauth2/token", conf.ServerUrl),
		EndpointParams: url.Values{
			"grant_type": []string{"client_credentials"},
		},
	}
	client := authConf.Client(ctx)
	var serverMode string
	if conf.ServerMode != "" {
		serverMode = fmt.Sprint("/", conf.ServerMode)
	}
	ver := version
	if conf.DebugEnabled {
		ver = fmt.Sprint(ver, "/debug")
	}
	return v1beta{
		requester: &requester{
			Client:  client,
			Url:     fmt.Sprint(conf.ServerUrl, serverMode),
			Version: ver,
		},
	}
}

func Wrap(ctx context.Context, client *http.Client, conf Config) Client {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)
	return New(ctx, conf)
}
