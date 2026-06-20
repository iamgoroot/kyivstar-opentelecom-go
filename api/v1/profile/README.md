# Profile — Profile API

Get subscriber profile information.

## Methods

| Method | HTTP | Path |
|--------|------|------|
| `Get` | POST | `/subscribers/profile` |

## Standalone Usage

```go
import (
	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/profile"
)

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := profile.NewService(ksClient)
resp, err := svc.Get(ctx, `{
 profile(msisdn:"380672000200"){
age
gender
 }
}`)
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.Profile.Get(ctx, query)
```
