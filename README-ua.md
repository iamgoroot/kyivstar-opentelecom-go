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
	"fmt"
	ksOpen "github.com/iamgoroot/kyivstar-opentelecom-go"
)

func main() {
   conf := ksOpen.Config{
       ServerUrl:   ksOpen.Gateway,
       ServerMode:  ksOpen.ServerModeMock,
       ClientID:    "your_client_id",
       ClientSecret: "your_client_secret",
   }
   ctx := context.Background()
   ksClient := ksOpen.New(ctx, conf)
   const destinationPhoneNumber = "380670000200"
	
   // Надсилання програмованих SMS
   sendMsgResp, err := ksClient.Send(
       ksOpen.SmsSendReq{
           From: "messagedesk",
           To:  destinationPhoneNumber,
           Text: "Hello World!",
       })
   fmt.Println("Sent", sendMsgResp, err)
   
   // Перевірка статусу програмованих SMS
   check, err := ksClient.Check(sendMsgResp.MsgId)
   fmt.Println("Check", check.Status, err)
   
   // Скоринг
   scoring, err := ksClient.Scoring(destinationPhoneNumber, 0)
   fmt.Println("Scored:", scoring, err)
   
   // Перевірка SIM
   sim, err := ksClient.VerifySim(destinationPhoneNumber, ksOpen.VerifySimReq{
       ActivationHours: 48,
   })
   fmt.Printf("Verify sim: changed=%d, active=%d, err=%v", sim.SimChanged, sim.IsActive, err)
}

```

## Документація

Для більш детальної інформації про використання клієнта, будь ласка, зверніться до [офіційної документації](https://api-gateway.kyivstar.ua).

## Участь

Контрибуції вітаються! Якщо ви знайдете помилку або маєте запит на функцію, будь ласка, відкрийте проблему (github issue) або надішліть PR.

## Ліцензія

Цей проект ліцензований за ліцензією MIT. Для отримання більш подробної інформації дивіться файл [LICENSE](LICENSE).

## Контакт

Якщо у вас є питання або вам потрібна допомога, будь ласка, відкрийте проблему (github issue) в цьому репозиторії або зверніться до підтримувачів безпосередньо.