# Viber — Viber Повідомлення

Send and check Viber messages (transactional and promotional).

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `SendTransaction` | POST | `/viber/transaction` |
| `SendPromotionText` | POST | `/viber/promotion` |
| `SendPromotionImage` | POST | `/viber/promotion` |
| `SendPromotionAction` | POST | `/viber/promotion` |
| `Check` | GET | `/viber/status/{msgId}` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/viber"

ksClient, _ := ksOpen.NewOauthClient(ctx, conf)
svc := viber.NewService(ksClient)
resp, err := svc.SendTransaction(ctx, viber.TransactionReq{From: "messagedesk", To: "380670000200", Text: "Hello!"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, conf)
ksClient.Viber.SendTransaction(ctx, req)
```

Each product can be used standalone via `product.NewService(client.Client{...})` or through the aggregated `V1Client` which bundles all products together.
