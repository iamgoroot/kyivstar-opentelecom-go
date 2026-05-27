# CLAUDE.md — Kyivstar Open Telecom API Go Client

## Project Overview

Unofficial Go client SDK for the [Kyivstar Open Telecom API](https://api-gateway.kyivstar.ua).
- **Module:** `github.com/iamgoroot/kyivstar-opentelecom-go`
- **Package:** `ksopentelecom`
- **Go version:** 1.26
- **Auth:** OAuth2 client-credentials (handled by `oauth.go`)
- **HTTP client:** `internal/client/requester.go` — generic `Get[Resp]` and `Post[Req, Resp]` wrappers

## Code Conventions

- Package name matches directory name (e.g., `api/v1/sms` → `package sms`)
- One product = one package under `api/v1/<product>/`
- Each product has exactly 3 files: `models.go`, `interface.go`, `client.go`
- Use `client.Client` for HTTP calls (never `http.Client` directly)
- Error types go in `internal/models/`; product-specific errors go in the product's `models.go`
- JSON tags on all exported struct fields
- No external code generation — write Go code by hand following the `api/v1/sms` template

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

#### `client.go` — Implementation using `client.Get` / `client.Post`

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
```

Use `client.Post[Req, Resp]` for POST endpoints, `client.Get[Resp]` for GET endpoints.
For URL query params, pass `url.Values` as the 4th argument.

### Step 3: Wire into `api.go`

```go
package ksopentelecom

import (
    "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/<product>"
    "github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type (
    SMS     = sms.Service
    <PRODUCT> = <product>.Service  // add
)

type V1Client struct {
    SMS
    <PRODUCT>  // add
}

func createV1Client(ksClient client.Client) (V1Client, error) {
    return V1Client{
        SMS:       sms.NewService(ksClient),
        <PRODUCT>: <product>.NewService(ksClient),  // add
    }, nil
}
```

### Step 4: Optional — create public request/response type aliases in root

If you want top-level convenience types (like `SmsSendReq`), add them in a file like `sms.go` (or directly in `api.go`):

```go
package ksopentelecom

import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"

type SmsSendReq = sms.SendReq
```

## How to Add a New API Version (e.g., v2)

1. Create `api/v2/` directory with product sub-packages (same 3-file pattern).
2. Create `v2.go` at the root:

```go
package ksopentelecom

import "context"

func NewV2Client(ctx context.Context, conf Config) (V2Client, error) {
    ksClient, err := createOauthClient(ctx, conf)
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
| RCS Повідомлення | `api/v1/rcs` | `/rcs/text`, `/rcs/suggestion`, `/rcs/richcard`, `/rcs/{msgId}` | POST, POST, POST, GET |
| Viber Повідомлення | `api/v1/viber` | `/viber/transaction`, `/viber/promotion`, `/viber/status/{msgId}` | POST, POST, GET |
| Розсилка на свій список | `api/v1/promo` | `/promo`, `/promo/{uuid}`, `/promo/{uuid}/audience`, `/promo/{uuid}/image`, `/promo/{uuid}/status/{status}`, `/promo/{uuid}/statistics` | POST, GET, GET, POST, POST, PUT, GET |
| Багатоканальна розсилка | `api/v1/multichannel` | `/messaging/multichannel`, `/messaging/multichannel/{multiMsgId}` | POST, GET |
| Sim Check Protection | `api/v1/simcheck` | `/subscribers/{phoneNumber}/sim-check-antifraud` | GET |
| Sim Count | `api/v1/simcount` | `/subscribers/{phoneNumber}/sim-count` | GET |
| Financial Scoring | `api/v1/scoring` | `/subscribers/{phoneNumber}/scoring` | GET |
| Lifetime Check | `api/v1/lifetime` | `/subscribers/{phoneNumber}/lifetime` | GET |
| Device Check | `api/v1/devicecheck` | `/subscribers/{phoneNumber}/device-check` | GET |
| SMS OTP Verification | `api/v1/otp` | `/verification/sms`, `/verification/sms/check` | POST, POST |
| Flash Call OTP | `api/v1/flashcall` | `/verification/flash-call`, `/verification/flash-call/check` | POST, POST |
| Profile API | `api/v1/profile` | `/subscribers/profile` | POST |

All `subscribers/*` paths use `{phoneNumber}` as a path parameter. Construct the endpoint path using `path.Join("v1/subscribers", phoneNumber, "scoring")`.

## Build & Test

```bash
go build ./...
go vet ./...
```

No test framework is currently configured; add `_test.go` files per package following Go conventions.
