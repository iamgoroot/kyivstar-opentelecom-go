package ksopentelecom

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ServerMode string

const (
	Gateway                      = "https://api-gateway.kyivstar.ua"
	ServerModeLive    ServerMode = ""
	ServerModeMock    ServerMode = "mock"
	ServerModeSandbox ServerMode = "sandbox"
)

type Config struct {
	ServerURL    string     `json:"serverUrl" yaml:"serverUrl"`
	ClientID     string     `json:"clientId" yaml:"clientId"`
	ClientSecret string     `json:"clientSecret" yaml:"clientSecret"`
	ServerMode   ServerMode `json:"serverMode" yaml:"serverMode"`
}

// LoadEnv overrides config fields from environment variables:
//
//	KS_CLIENT_ID       — OAuth2 client ID
//	KS_CLIENT_SECRET   — OAuth2 client secret
//	KS_SERVER_URL      — API gateway base URL
//	KS_SERVER_MODE     — Server mode: "mock", "sandbox", or "live"
func (c *Config) LoadEnv() error {
	c.ClientID = cmp.Or(os.Getenv("KS_CLIENT_ID"), c.ClientID)
	c.ClientSecret = cmp.Or(os.Getenv("KS_CLIENT_SECRET"), c.ClientSecret)
	c.ServerURL = cmp.Or(os.Getenv("KS_SERVER_URL"), c.ServerURL)

	if v := os.Getenv("KS_SERVER_MODE"); v != "" {
		switch v {
		case "mock":
			c.ServerMode = ServerModeMock
		case "sandbox":
			c.ServerMode = ServerModeSandbox
		case "live":
			c.ServerMode = ServerModeLive
		default:
			return fmt.Errorf("invalid value for KS_SERVER_MODE: %s", v)
		}
	}

	return nil
}

func (c *Config) LoadJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}
