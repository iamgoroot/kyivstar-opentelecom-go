# CLAUDE.md — Kyivstar Open Telecom API Go Client

## Project Overview

Unofficial Go client SDK for the [Kyivstar Open Telecom API](https://api-gateway.kyivstar.ua).
- **Module:** `github.com/iamgoroot/kyivstar-opentelecom-go`
- **Package:** `ksopentelecom`
- **Go version:** 1.26
- **Auth:** OAuth2 client-credentials (handled by `oauth.go`)
- **HTTP client:** `internal/client/requester.go` — generic `Get[Resp]`, `Post[Req, Resp]`, and `Put[Req, Resp]` wrappers

## Code Conventions

- Package name matches directory name (e.g., `api/v1/sms` → `package sms`)
- One product = one package under `api/v1/<product>/`
- Each product has exactly 4 files: `models.go`, `interface.go`, `client.go`, `README.md`
- Use `client.Client` for HTTP calls (never `http.Client` directly)
- Error types go in `internal/models/`; product-specific errors go in the product's `models.go`
- JSON tags on all exported struct fields
- No external code generation — write Go code by hand following the `api/v1/sms` template
- Use `new("value")` / `new(42)` to create pointers to literal values (Go 1.26+) instead of writing helper functions
- Treat `openapi.yaml` as source of truth for all endpoints, request/response schemas, and product structure. The OpenAPI spec fetched from `https://api-gateway.kyivstar.ua/api/public/openapi.yaml` overrides any secondary documentation.

## How to Add a New API Product

### Step 1: Read the OpenAPI spec

Fetch the spec:
```
https://api-gateway.kyivstar.ua/api/public/openapi.yaml
```

Identify the tag (product), its paths, methods, and request/response schemas.

### Step 2: Create the product package

Create `api/v1/<product>/` with 3 files.

#### `models.go` — Request/response structs

```go
package <product>

type SomeReq struct {
    Field string `json:"field"`
}

type SomeResp struct {
    ReqID string `json:"reqId"`
    Field string `json:"field"`
}
```

#### `interface.go` — Service interface + constructor

```go
package <product>

import (
    "context"
    "github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type Service interface {
    DoSomething(ctx context.Context, req SomeReq) (SomeResp, error)
    GetSomething(ctx context.Context, id string) (SomeResp, error)
}

func NewService(client client.Client) Service {
    return &service{client: client}
}
```

#### `client.go` — Implementation using `client.Get` / `client.Post` / `client.Put`

```go
package <product>

import (
    "context"
    "path"
    "github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

const endpointContextPath = "v1/<product>"

type service struct {
    client client.Client
}

func (s service) DoSomething(ctx context.Context, req SomeReq) (SomeResp, error) {
    return client.Post[SomeReq, SomeResp](ctx, s.client, endpointContextPath, nil, req)
}

func (s service) GetSomething(ctx context.Context, id string) (SomeResp, error) {
    endpointPath := path.Join(endpointContextPath, id)
    return client.Get[SomeResp](ctx, s.client, endpointPath, nil)
}

func (s service) UpdateSomething(ctx context.Context, id string) (SomeResp, error) {
    endpointPath := path.Join(endpointContextPath, id)
    return client.Put[struct{}, SomeResp](ctx, s.client, endpointPath, nil, struct{}{})
}
```

Use `client.Post[Req, Resp]` for POST, `client.Get[Resp]` for GET, `client.Put[Req, Resp]` for PUT.
For URL query params, pass `url.Values` as the 4th argument.

### Step 3: Wire into `api.go`

```go
package ksopentelecom

import (
    "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/<product>"
    "github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type V1Client struct {
    SMS       sms.Service
    <PRODUCT> <product>.Service  // add
}

func createV1Client(ksClient client.Client) (V1Client, error) {
    return V1Client{
        SMS:       sms.NewService(ksClient),
        <PRODUCT>: <product>.NewService(ksClient),  // add
    }, nil
}
```

Note: `V1Client` uses **plain named fields** (not embedding). Access services explicitly:
```go
ksClient.SMS.Send(ctx, req)
ksClient.Scoring.Check(ctx, phone, formula)
```

### Step 4: Optional — create public request/response type aliases in root

If you want top-level convenience types (like `SmsSendReq`), add them in a file like `sms.go` (or directly in `api.go`):

```go
package ksopentelecom

import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"

type SmsSendReq = sms.SendReq
```

### Step 5: Validate completeness
After implementing, verify that the product implementation covers **all endpoints** listed in the OpenAPI spec for that tag (product). Compare the methods in your `interface.go` against each path and HTTP method under that tag in `openapi.yaml`. Every endpoint must have a corresponding method in the `Service` interface.

**Exclude Webhooks** from client implementation — it is a server-side callback feature, not a client API product. No `api/v1/webhooks/` package should be created.

## How to Add a New API Version (e.g., v2)

1. Create `api/v2/` directory with product sub-packages (same 3-file pattern).
2. Create `v2.go` at the root:

```go
package ksopentelecom

import "context"

func NewV2Client(ctx context.Context, conf Config) (V2Client, error) {
    ksClient, err := NewOauthClient(ctx, conf)
    if err != nil {
        return V2Client{}, err
    }
    return createV2Client(ksClient)
}
```

3. Create the `V2Client` struct in `api.go` or a dedicated `api_v2.go`.

## All API Products (from OpenAPI spec)

| Product | Package | Paths | HTTP Methods |
|---------|---------|-------|-------------|
| Програмовані SMS | `api/v1/sms` ✅ done | `/sms`, `/sms/{msgId}`, `/sms/batch`, `/sms/status/batch` | POST, GET, POST, POST |
| RCS Повідомлення | `api/v1/rcs` ✅ done | `/rcs/text`, `/rcs/suggestion`, `/rcs/richcard`, `/rcs/{msgId}` | POST, POST, POST, GET |
| Viber Повідомлення | `api/v1/viber` ✅ done | `/viber/transaction`, `/viber/promotion`, `/viber/status/{msgId}` | POST, POST, GET |
| Розсилка на свій список | `api/v1/promo` ✅ done | `/promo`, `/promo/{uuid}`, `/promo/{uuid}/audience`, `/promo/{uuid}/image`, `/promo/{uuid}/status/{status}`, `/promo/{uuid}/statistics` | POST, GET, GET, POST, POST, PUT, GET |
| Багатоканальна розсилка | `api/v1/multichannel` ✅ done | `/messaging/multichannel`, `/messaging/multichannel/{multiMsgId}` | POST, GET |
| Sim Check Protection | `api/v1/simcheck` ✅ done | `/subscribers/{phoneNumber}/sim-check-antifraud` | GET |
| Sim Count | `api/v1/simcount` ✅ done | `/subscribers/{phoneNumber}/sim-count` | GET |
| Financial Scoring | `api/v1/scoring` ✅ done | `/subscribers/{phoneNumber}/scoring` | GET |
| Lifetime Check | `api/v1/lifetime` ✅ done | `/subscribers/{phoneNumber}/lifetime` | GET |
| Device Check | `api/v1/devicecheck` ✅ done | `/subscribers/{phoneNumber}/device-check` | GET |
| SMS OTP Verification | `api/v1/otp` ✅ done | `/verification/sms`, `/verification/sms/check` | POST, POST |
| Flash Call OTP | `api/v1/flashcall` ✅ done | `/verification/flash-call`, `/verification/flash-call/check` | POST, POST |
| Profile API | `api/v1/profile` ✅ done | `/subscribers/profile` | POST |

All `subscribers/*` paths use `{phoneNumber}` as a path parameter. Construct the endpoint path using `path.Join("v1/subscribers", phoneNumber, "scoring")`.

## Examples

Standalone examples live in `examples/<product>/` — each has its own `go.mod` and can be built independently.

```
examples/
├── sms/           # sms.NewService(ksClient) — standalone usage
├── rcs/           # rcs.NewService(ksClient)
├── viber/
├── promo/
├── multichannel/
├── simcheck/
├── simcount/
├── scoring/
├── lifetime/
├── devicecheck/
├── otp/
├── flashcall/
├── profile/
└── all/           # ksOpen.NewV1Client(ctx, conf) — aggregated V1Client usage
```

Each product example uses `NewOauthClient` for auth and `product.NewService(ksClient)` directly. The `all` example demonstrates `V1Client` which bundles all products.

Create new examples at `examples/<product>/main.go` when adding a new product.

## Public API

| Function | Purpose |
|----------|---------|
| `NewOauthClient(ctx, conf)` | Creates authenticated `client.Client` for direct product usage |
| `NewV1Client(ctx, conf)` | Creates `V1Client` with all products wired together |

## Integration Tests

Tests live in `test/local/` (same Go module, no separate `go.mod`).

### Structure

```
test/
└── local/
    ├── handlers/
    │   ├── server.go           # NewServer(register funcs) + writeJSON helper
    │   ├── sms.go              # RegisterSMS(mux *http.ServeMux)
    │   ├── rcs.go              # RegisterRCS(mux)
    │   └── ...                 # one per product
    ├── sms_test.go             # tests for sms.Service
    ├── rcs_test.go             # tests for rcs.Service
    └── ...                     # one per product
```

### How to add a handler

Create `test/local/handlers/<product>.go`:

```go
package handlers

import "net/http"

func RegisterSMS(mux *http.ServeMux) {
    mux.HandleFunc("POST /rest/v1/sms", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, map[string]any{
            "reqId": "ad30594292f7959683a410bf1add088e",
            "msgId": "20200000-0000-0000-0000-380670000200",
        })
    })
    mux.HandleFunc("GET /rest/v1/sms/{msgId}", func(w http.ResponseWriter, r *http.Request) {
        writeJSON(w, map[string]interface{}{
            "reqId":  "ad30594292f7959683a410bf1add088e",
            "msgId":  r.PathValue("msgId"),
            "status": "delivered",
        })
    })
}
```

Use hardcoded example values from the OpenAPI spec.

### How to add a test

Create `test/local/<product>_test.go`:

```go
package testinglocal

import (
    "context"
    "testing"
    "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
    "github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
    "github.com/iamgoroot/kyivstar-opentelecom-go/test/local/handlers"
)

func TestSMSSend(t *testing.T) {
    srv := handlers.NewServer(handlers.RegisterSMS)
    defer srv.Close()

    svc := sms.NewService(client.Client{Client: srv.Client(), BaseUrl: srv.URL})
    resp, err := svc.Send(context.Background(), sms.SendReq{...})
    if err != nil {
        t.Fatal(err)
    }
    if resp.MsgID != "20200000-0000-0000-0000-380670000200" {
        t.Errorf("unexpected msgID: %s", resp.MsgID)
    }
}
```

### Running tests

```bash
go test ./test/local/... -v
```

## Build & Test

```bash
go build ./...
go vet ./...
go vulncheck ./...
golangci-lint run ./...
go test ./test/local/... -v
```

Always use `go build -C <dir> -o /dev/null` for examples to avoid leaving compiled binaries: `-C` must come before `-o`:
```bash
go build -C examples/sms -o /dev/null .
go build -C examples/all -o /dev/null .
```
