# Multichannel — Багатоканальна розсилка

Send messages across multiple channels and check their status.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Send` | POST | `/messaging/multichannel` |
| `Check` | GET | `/messaging/multichannel/{multiMsgId}` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/multichannel"

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := multichannel.NewService(ksClient)
resp, err := svc.Send(ctx, multichannel.SendReq{
	To: "380670000200",
	SmsContent: &multichannel.SmsContent{
		From: "messagedesk",
		Text: "Hello!",
	},
})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.Multichannel.Send(ctx, req)
```
