# SMS — Програмовані SMS

Send and check programmable SMS messages.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Send` | POST | `/sms` |
| `SendBatch` | POST | `/sms/batch` |
| `Check` | GET | `/sms/{msgId}` |
| `CheckBatch` | POST | `/sms/status/batch` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := sms.NewService(ksClient)
resp, err := svc.Send(ctx, sms.SendReq{From: "messagedesk", To: "380670000200", Text: "Hello!"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.SMS.Send(ctx, req)
```
