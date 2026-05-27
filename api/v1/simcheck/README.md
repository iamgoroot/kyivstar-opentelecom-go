# SimCheck — Sim Check Protection

Check if a subscriber's SIM card has been recently changed (anti-fraud).

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Check` | GET | `/subscribers/{phoneNumber}/sim-check-antifraud` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcheck"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := simcheck.NewService(ksClient)
resp, err := svc.Check(ctx, "380670000200", 7)
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.SimCheck.Check(ctx, phone, period)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
