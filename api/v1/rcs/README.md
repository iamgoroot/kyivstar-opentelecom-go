# RCS — RCS Повідомлення

Send and check Rich Communication Services messages.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `SendText` | POST | `/rcs/text` |
| `SendSuggestion` | POST | `/rcs/suggestion` |
| `SendRichCard` | POST | `/rcs/richcard` |
| `Check` | GET | `/rcs/{msgId}` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/rcs"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := rcs.NewService(ksClient)
resp, err := svc.SendText(ctx, rcs.RcsTextReq{From: "messagedesk", To: "380670000200", Text: "Hello!"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.RCS.SendText(ctx, req)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
