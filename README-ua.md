# Kyivstar Open Telecom API Клієнт для Go

Це неофіційний клієнт Go для API Kyivstar, який дозволяє вам взаємодіяти з API Київстар відкритий телеком https://api-gateway.kyivstar.ua

## Встановлення

Щоб встановити API клієнт для Київстар відкритий телеком використовуйте команду `go get`:

```bash
go get github.com/iamgoroot/kyivstar-opentelecom-go
```
## Перед початком

Щоб використовувати API, вам потрібно зареєструватися на https://api-market.kyivstar.ua та отримати client_id та client_secret.

## Використання

Ось простий приклад використання клієнта:

```go
package main

import (
	"context"
	"log"

	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
	"github.com/iamgoroot/kyivstar-opentelecom-go/api/v1/sms"
)

func main() {
	var conf ksOpen.Config
	if err := conf.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ksClient, err := ksOpen.NewV1Client(ctx, &conf)
	if err != nil {
		log.Fatal(err)
	}

	const destinationPhoneNumber = "380670000200"

	// Надсилання програмованих SMS
	sendMsgResp, err := ksClient.SMS.Send(
		ctx,
		sms.SendReq{
			From: "messagedesk",
			To:   destinationPhoneNumber,
			Text: "Hello World!",
		},
	)
	log.Println("Sent", sendMsgResp, err)

	// Перевірка статусу програмованих SMS
	check, err := ksClient.SMS.Check(ctx, sendMsgResp.MsgID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Check", check.Status, err)
}
```

## Документація

Для більш детальної інформації про використання клієнта зверніться до [офіційної документації](https://api-gateway.kyivstar.ua).

## Участь

Контрибуції вітаються! Якщо ви знайдете помилку або маєте пропозицію щодо нової функції, будь ласка, відкрийте проблему (github issue) або надішліть PR.

## Ліцензія

Цей проект ліцензований за ліцензією MIT. Для отримання більш детальної інформації дивіться файл [LICENSE](LICENSE).
