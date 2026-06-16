package testinglocal

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
	"github.com/iamgoroot/kyivstar-opentelecom-go/test/handlers"
)

func retryOnRateLimit(t *testing.T, fn func() error) {
	t.Helper()

	for range 10 {
		err := fn()
		if err == nil {
			return
		}

		var ke models.KotError
		if errors.As(err, &ke) && ke.HTTPStatus == http.StatusTooManyRequests {
			t.Logf("rate limited, waiting %ds...", ke.Info.Reset)
			time.Sleep(time.Duration(ke.Info.Reset)*time.Second + 100*time.Millisecond)

			continue
		}

		t.Fatal(err)
	}

	t.Fatal("max retries exceeded")
}

func asKotError(err error, out *models.KotError) bool {
	if ke, ok := errors.AsType[models.KotError](err); ok {
		*out = ke
		return true
	}

	return false
}

func setupTestClient(t requireTestT, registers ...func(*http.ServeMux)) client.Client {
	t.Helper()

	ctx := context.Background()

	conf := ksOpen.Config{
		ServerURL:    os.Getenv("KS_SERVER_URL"),
		ClientID:     os.Getenv("KS_CLIENT_ID"),
		ClientSecret: os.Getenv("KS_CLIENT_SECRET"),
	}
	if err := conf.LoadEnv(); err != nil {
		t.Fatalf("LoadEnv: %v", err)
	}

	if conf.ClientID != "" && conf.ClientSecret != "" && conf.ServerURL != "" {
		ksClient, err := ksOpen.NewOauthClient(ctx, &conf)
		if err != nil {
			t.Fatalf("NewOauthClient: %v", err)
		}

		return ksClient
	}

	srv := handlers.NewServer(registers...)
	t.Cleanup(srv.Close)

	return client.Client{Client: srv.Client(), BaseURL: srv.URL}
}

type requireTestT interface {
	Helper()
	Fatalf(format string, args ...any)
	Cleanup(func())
}
