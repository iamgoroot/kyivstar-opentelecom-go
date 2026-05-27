# Profile — Profile API

Get subscriber profile information.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Get` | POST | `/subscribers/profile` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := profile.NewService(ksClient)
resp, err := svc.Get(ctx, "380670000200")
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.Profile.Get(ctx, query)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
