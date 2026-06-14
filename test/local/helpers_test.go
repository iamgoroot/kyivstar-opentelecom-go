package testinglocal

import (
	"context"
	"errors"
	"net/http"
	"os"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func asKotError(err error, out *models.KotError) bool {
	if ke, ok := errors.AsType[models.KotError](err); ok {
		*out = ke
		return true
	}

	return false
}

func isRunningLocally() bool {
	return os.Getenv("KS_CLIENT_ID") == "" || os.Getenv("KS_CLIENT_SECRET") == "" || os.Getenv("KS_SERVER_URL") == ""
}

func setupTestClient(t requireTestT, registers ...func(*http.ServeMux)) client.Client {
	t.Helper()

	clientID := os.Getenv("KS_CLIENT_ID")
	clientSecret := os.Getenv("KS_CLIENT_SECRET")
	serverURL := os.Getenv("KS_SERVER_URL")
	serverMode := os.Getenv("KS_SERVER_MODE")

	if clientID != "" && clientSecret != "" && serverURL != "" {
		ctx := context.Background()
		conf := ksOpen.Config{
			ServerUrl:    serverURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}

		switch serverMode {
		case "mock":
			conf.ServerMode = ksOpen.ServerModeMock
		case "sandbox":
			conf.ServerMode = ksOpen.ServerModeSandbox
		case "live", "":
			conf.ServerMode = ksOpen.ServerModeLive
		}

		ksClient, err := ksOpen.NewOauthClient(ctx, conf)
		if err != nil {
			t.Fatalf("NewOauthClient: %v", err)
		}

		return ksClient
	}

	srv := handlers.NewServer(registers...)
	t.Cleanup(srv.Close)

	return client.Client{Client: srv.Client(), BaseUrl: srv.URL}
}

type requireTestT interface {
	Helper()
	Fatalf(format string, args ...any)
	Cleanup(func())
}
