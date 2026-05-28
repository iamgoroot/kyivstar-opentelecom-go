# Kyivstar Open Telecom API Client for Go

Unofficial Go client SDK for the [Kyivstar Open Telecom API](https://api-gateway.kyivstar.ua). Provides full coverage of all API products — SMS, RCS, Viber, promo campaigns, multichannel messaging, SIM security checks, financial scoring, OTP verification, and more.

```go
import ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
```

## Installation

```bash
go get github.com/iamgoroot/kyivstar-opentelecom-go
```

## Before you start

1. Register at https://api-market.kyivstar.ua
2. Obtain your `client_id` and `client_secret`
3. Choose a server mode:
   - `ksOpen.ServerModeMock` — fake data, no charges, no real delivery
   - `ksOpen.ServerModeSandbox` — real data, limited to OTP-verified phone numbers, no charges
   - `ksOpen.ServerModeLive` — production, real delivery, charged

## Usage

Two approaches are available:

### 1. V1Client — all products, one client

The `V1Client` bundles every product service into a single struct. Use it when you need multiple API products in your application.

```go
package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	conf := ksOpen.Config{
		ServerUrl:    ksOpen.Gateway,
		ServerMode:   ksOpen.ServerModeMock,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	}

	ctx := context.Background()

	ksClient, err := ksOpen.NewV1Client(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	// Send SMS
	sendResp, err := ksClient.SMS.Send(ctx, sms.SendReq{
		From: "messagedesk",
		To:   "380670000200",
		Text: "Hello from Kyivstar API!",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SMS sent: msgID=%s", sendResp.MsgID)

	// Check SMS status
	checkResp, err := ksClient.SMS.Check(ctx, sendResp.MsgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SMS status: %s (updated %v)", checkResp.Status, checkResp.Date)
}
```

### 2. Standalone service — one product at a time

Use `NewOauthClient` + product-specific `NewService` when you only need a single API product.

```go
package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	ctx := context.Background()

	ksClient, err := ksOpen.NewOauthClient(ctx, ksOpen.Config{
		ServerUrl:    ksOpen.Gateway,
		ServerMode:   ksOpen.ServerModeMock,
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := sms.NewService(ksClient)

	resp, err := svc.Send(ctx, sms.SendReq{
		From: "messagedesk",
		To:   "380670000200",
		Text: "Hello World!",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SMS sent: reqID=%s msgID=%s", resp.ReqID, resp.MsgID)

	check, err := svc.Check(ctx, resp.MsgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SMS status: %s", check.Status)
}
```

## Available API products

| Product | Package | Methods |
|---------|---------|---------|
| SMS (programmable) | [`api/v1/sms`](api/v1/sms/README.md) | `Send`, `SendBatch`, `Check`, `CheckBatch` |
| RCS messaging | [`api/v1/rcs`](api/v1/rcs/README.md) | `SendText`, `SendSuggestion`, `SendRichCard`, `Check` |
| Viber messaging | [`api/v1/viber`](api/v1/viber/README.md) | `SendTransaction`, `SendPromotionText`, `SendPromotionImage`, `SendPromotionAction`, `Check` |
| Promo campaigns | [`api/v1/promo`](api/v1/promo/README.md) | `CreateSMS`, `CreateViber`, `CreateRCS`, `List`, `Get`, `AddAudience`, `AddImage`, `ChangeStatus`, `GetStatistics` |
| Multichannel messaging | [`api/v1/multichannel`](api/v1/multichannel/README.md) | `Send`, `Check` |
| SIM check protection | [`api/v1/simcheck`](api/v1/simcheck/README.md) | `Check(phone, period)` |
| SIM count check | [`api/v1/simcount`](api/v1/simcount/README.md) | `Check(phone, days)` |
| Financial scoring | [`api/v1/scoring`](api/v1/scoring/README.md) | `Check(phone, formula)` |
| Lifetime check | [`api/v1/lifetime`](api/v1/lifetime/README.md) | `Check(phone)` |
| Device check | [`api/v1/devicecheck`](api/v1/devicecheck/README.md) | `Check(phone)`, `CheckWithImei(phone, imei, days)` |
| SMS OTP verification | [`api/v1/otp`](api/v1/otp/README.md) | `Send`, `Check` |
| Flash call OTP | [`api/v1/flashcall`](api/v1/flashcall/README.md) | `Create`, `Check` |
| Profile API | [`api/v1/profile`](api/v1/profile/README.md) | `Get(query)` |

## Error handling

The client returns Go errors for HTTP-level failures (network, auth, 4xx/5xx). Product-specific error codes like `1001003` (unsupported alpha name) are returned as fields in the response struct — they are not converted to Go errors, so you should check `errorCode` / `errorMsg` fields where applicable.

Example:

```go
resp, err := svc.Send(ctx, req)
if err != nil {
    // network, auth, or 5xx error
    log.Fatal(err)
}
// Some endpoints return partial errors in the response body
if resp.ErrorCode != "" {
    log.Printf("API error: %s — %s", resp.ErrorCode, resp.ErrorMsg)
}
```

## Configuration

| Constant | Value | Description |
|----------|-------|-------------|
| `ksOpen.Gateway` | `https://api-gateway.kyivstar.ua` | Base URL |
| `ServerModeLive` | `""` | Production server |
| `ServerModeMock` | `"mock"` | Mock server (fake data, free) |
| `ServerModeSandbox` | `"sandbox"` | Sandbox (real data, limited) |

All three server modes — `Mock`, `Sandbox`, and `Live` — are supported. The `Config.ServerMode` field is appended to the base URL:

```
https://api-gateway.kyivstar.ua/rest/v1/…           ← Live
https://api-gateway.kyivstar.ua/mock/rest/v1/…      ← Mock
https://api-gateway.kyivstar.ua/sandbox/rest/v1/…   ← Sandbox
```

## Documentation

- [Official API documentation](https://api-gateway.kyivstar.ua)
- [OpenAPI specification](https://api-gateway.kyivstar.ua/api/public/openapi.yaml)
- Product-specific README files are in each package under `api/v1/<product>/`

## Examples

Standalone runnable examples for every product are in the `examples/` directory:

```bash
# Run the V1Client example (all products bundled)
go run -C examples/all .

# Run a standalone product example
go run -C examples/sms .
go run -C examples/scoring .
go run -C examples/otp .
```

## Contributing

Contributions are welcome! See [CLAUDE.md](./CLAUDE.md) for the contributor guide, coding conventions, and instructions on adding new API products.

## License

MIT — see [LICENSE](./LICENSE).

## Suggestions for further README improvements

1. **Quick-start guide** — a step-by-step tutorial that walks through registering on the portal, obtaining credentials, and making the first API call in under 5 minutes.
2. **FAQ section** — common issues: rate limits, alpha-name registration, unsupported phone number formats, TTL behaviour, batch size limits.
3. **Webhook handling** — although webhooks are server-side and not part of this client, a short section explaining how to validate and parse incoming webhook payloads would help users set up the receiving end.
4. **Comparison table** — when to use V1Client vs standalone services vs direct `client.Client`.
5. **Advanced auth section** — token refresh strategies, handling 401 responses with automatic retry, custom `http.Client` injection (e.g. with tracing or custom timeouts).
6. **Migration guide** — if the upstream API introduces breaking changes, document how to migrate between client versions.
7. **Performance section** — connection pooling, batch size recommendations, concurrency patterns.
8. **SDK reference** — auto-generated Go doc links (pkg.go.dev) once the module is published.
