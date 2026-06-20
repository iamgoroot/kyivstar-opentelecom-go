# Promo — Розсилка на свій список

Create and manage bulk SMS, Viber, and RCS campaigns to your own subscriber list.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `CreateSMS` | POST | `/promo` |
| `CreateViber` | POST | `/promo` |
| `CreateRCS` | POST | `/promo` |
| `List` | GET | `/promo` |
| `Get` | GET | `/promo/{uuid}` |
| `AddAudience` | POST | `/promo/{uuid}/audience` |
| `AddImage` | POST | `/promo/{uuid}/image` |
| `ChangeStatus` | PUT | `/promo/{uuid}/status/{status}` |
| `GetStatistics` | GET | `/promo/{uuid}/statistics` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/promo"

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := promo.NewService(ksClient)
p, err := svc.CreateSMS(ctx, promo.CreateSMSReq{From: "author", Text: "Hello ${1}", CampaignType: "SMS"})
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.Promo.CreateSMS(ctx, req)
```
