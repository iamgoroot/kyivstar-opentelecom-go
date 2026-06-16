# DeviceCheck — Device Check

Get device information for a subscriber.

> **Note:** The official Kyivstar Open Telecom API documentation for Device Check is incorrect.
> The documented method signatures do not match the actual API. The correct signatures are:

| Method | HTTP | Path | Query Params |
|--------|------|------|-------------|
| `Check(ctx, phoneNumber, imei)` | GET | `/subscribers/{phoneNumber}/device-check` | `imei` |
| `CheckWithImei(ctx, phoneNumber, daysPeriod)` | GET | `/subscribers/{phoneNumber}/device-check` | `daysPeriod` |

## Standalone Usage

```go
import (
	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/devicecheck"
)

ksClient, _ := ksOpen.NewOauthClient(ctx, &conf)
svc := devicecheck.NewService(ksClient)

// Check with IMEI
resp, err := svc.Check(ctx, "380670000200", "123456789012345")

// Check with days period
_, err = svc.CheckWithImei(ctx, "380670000200", 30)
```

## Aggregated Usage (V1Client)

```go
ksClient, _ := ksOpen.NewV1Client(ctx, &conf)
ksClient.DeviceCheck.Check(ctx, phone, imei)
ksClient.DeviceCheck.CheckWithImei(ctx, phone, 30)
```
