# FlashCall — Flash Call OTP

Verify phone numbers via flash call (missed call with OTP in caller ID).

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Create` | POST | `/verification/flash-call` |
| `Check` | POST | `/verification/flash-call/check` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/flashcall"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := flashcall.NewService(ksClient)
resp, err := svc.Create(ctx, flashcall.CreateReq{PhoneNumber: "380670000200"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.FlashCall.Create(ctx, req)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
