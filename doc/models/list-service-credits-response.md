
# List Service Credits Response

## Structure

`ListServiceCreditsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceCredits` | [`[]models.ServiceCredit1`](../../doc/models/service-credit-1.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listServiceCreditsResponse := models.ListServiceCreditsResponse{
        ServiceCredits:       []models.ServiceCredit1{
            models.ServiceCredit1{
                Id:                      models.ToPointer(224),
                AmountInCents:           models.ToPointer(int64(54)),
                EndingBalanceInCents:    models.ToPointer(int64(94)),
                EntryType:               models.ToPointer(models.ServiceCreditType_CREDIT),
                Memo:                    models.ToPointer("memo2"),
            },
            models.ServiceCredit1{
                Id:                      models.ToPointer(224),
                AmountInCents:           models.ToPointer(int64(54)),
                EndingBalanceInCents:    models.ToPointer(int64(94)),
                EntryType:               models.ToPointer(models.ServiceCreditType_CREDIT),
                Memo:                    models.ToPointer("memo2"),
            },
            models.ServiceCredit1{
                Id:                      models.ToPointer(224),
                AmountInCents:           models.ToPointer(int64(54)),
                EndingBalanceInCents:    models.ToPointer(int64(94)),
                EntryType:               models.ToPointer(models.ServiceCreditType_CREDIT),
                Memo:                    models.ToPointer("memo2"),
            },
        },
    }

}
```

