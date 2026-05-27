# DeviceCheck — Device Check

Get device information for a subscriber.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Check` | GET | `/subscribers/{phoneNumber}/device-check` |
| `CheckWithImei` | GET | `/subscribers/{phoneNumber}/device-check` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := devicecheck.NewService(ksClient)
resp, err := svc.Check(ctx, "380670000200")
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.DeviceCheck.Check(ctx, phone)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
