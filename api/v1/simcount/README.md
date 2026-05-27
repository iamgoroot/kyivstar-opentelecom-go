# SimCount — Sim Count

Get the number of SIM cards registered to a subscriber.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Check` | GET | `/subscribers/{phoneNumber}/sim-count` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/simcount"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := simcount.NewService(ksClient)
resp, err := svc.Check(ctx, "380670000200", 30)
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.SimCount.Check(ctx, phone, days)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
