package ksopentelecom

type serverMode string

const (
	Gateway                      = "https://api-gateway.kyivstar.ua"
	ServerModeLive    serverMode = ""
	ServerModeMock    serverMode = "mock"
	ServerModeSandbox serverMode = "sandbox"
)

type Config struct {
	ServerUrl    string     `json:"serverUrl" yaml:"serverUrl"`
	ClientID     string     `json:"clientId" yaml:"clientId"`
	ClientSecret string     `json:"clientSecret" yaml:"clientSecret"`
	ServerMode   serverMode `json:"serverMode" yaml:"serverMode"`
}
