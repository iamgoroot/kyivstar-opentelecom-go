# Scoring — Financial Scoring

Perform financial scoring of a subscriber.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Check` | GET | `/subscribers/{phoneNumber}/scoring` |

## Standalone Usage

```go
import "github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/scoring"

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := scoring.NewService(ksClient)
resp, err := svc.Check(ctx, "380670000200", 1)
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.Scoring.Check(ctx, phone, formula)
```
