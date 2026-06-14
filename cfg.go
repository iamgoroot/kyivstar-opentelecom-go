package ksopentelecom

type ServerMode string

const (
	Gateway                      = "https://api-gateway.kyivstar.ua"
	ServerModeLive    ServerMode = ""
	ServerModeMock    ServerMode = "mock"
	ServerModeSandbox ServerMode = "sandbox"
)

type Config struct {
	ServerUrl    string     `json:"serverUrl" yaml:"serverUrl"`
	ClientID     string     `json:"clientId" yaml:"clientId"`
	ClientSecret string     `json:"clientSecret" yaml:"clientSecret"`
	ServerMode   ServerMode `json:"serverMode" yaml:"serverMode"`
}
