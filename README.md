# Kyivstar Open Telecom API Client for Go

This is unofficial a Go client for the Kyivstar API, which allows you to interact with the Kyivstar Open Telecom API https://api-gateway.kyivstar.ua

## Installation

To install the Kyivstar API client, use the `go get` command:

```bash
go get github.com/iamgoroot/kyivstar-opentelecom-go
```
## Before you start

In order to use the API, you need to register at https://api-market.kyivstar.ua and obtain client_id and client_secret.

## Usage

Here's a simple example of how to use the Kyivstar API client:

```go

package main

import (
	"context"
	"fmt"
	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
)

func main() {
    conf := ksOpen.Config{
        ServerUrl:    ksOpen.Gateway,
        ServerMode:   ksOpen.ServerModeMock,
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    }
    ctx := context.Background()
    ksClient := ksOpen.New(ctx, conf)
    const destinationPhoneNumber = "380670000200"
	
    // Send SMS msg
    sendMsgResp, err := ksClient.Send(
        ksOpen.SmsSendReq{
            From: "messagedesk",
            To:   destinationPhoneNumber,
            Text: "Hello World!",
        })
    fmt.Println("Sent", sendMsgResp, err)
    
    // Check SMS Status
    check, err := ksClient.Check(sendMsgResp.MsgId)
    fmt.Println("Check", check.Status, err)
    
    // Scoring
    scoring, err := ksClient.Scoring(destinationPhoneNumber, 0)
    fmt.Println("Scored:", scoring, err)
    
    // Verify sim
    sim, err := ksClient.VerifySim(destinationPhoneNumber, ksOpen.VerifySimReq{
        ActivationHours: 48,
    })
    fmt.Printf("Verify sim: changed=%d, active=%d, err=%v", sim.SimChanged, sim.IsActive, err)
}

```

## Documentation

For more detailed information on how to use the client, please refer to the [official documentation](https://api-gateway.kyivstar.ua).

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

If you have any questions or need assistance, please open an issue on this repository or contact the maintainers directly.

