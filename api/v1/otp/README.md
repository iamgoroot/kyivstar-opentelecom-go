# OTP — SMS OTP Verification

Send and verify one-time passwords via SMS.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Send` | POST | `/verification/sms` |
| `Check` | POST | `/verification/sms/check` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/otp"

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := otp.NewService(ksClient)
resp, err := svc.Send(ctx, otp.SendReq{To: "380670000200"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.OTP.Send(ctx, req)
```
