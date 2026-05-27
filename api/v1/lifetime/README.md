# Lifetime — Lifetime Check

Get the lifetime (in days) of a subscriber's number.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Check` | GET | `/subscribers/{phoneNumber}/lifetime` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/lifetime"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := lifetime.NewService(ksClient)
resp, err := svc.Check(ctx, "380670000200")
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.Lifetime.Check(ctx, phone)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
