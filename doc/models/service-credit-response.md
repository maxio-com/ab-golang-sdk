
# Service Credit Response

## Structure

`ServiceCreditResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceCredit` | [`models.ServiceCredit`](../../doc/models/service-credit.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    serviceCreditResponse := models.ServiceCreditResponse{
        ServiceCredit:        models.ServiceCredit{
            Id:                   models.ToPointer(38),
            AmountInCents:        models.ToPointer(int64(124)),
            EndingBalanceInCents: models.ToPointer(int64(164)),
            EntryType:            models.ToPointer(models.ServiceCreditType_CREDIT),
            Memo:                 models.ToPointer("memo0"),
        },
    }

}
```

