
# Subscription Group Prepayment Response

## Structure

`SubscriptionGroupPrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry. |
| `EndingBalanceInCents` | `*int64` | Optional | The ending balance in cents of the account. |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `models.Optional[string]` | Optional | A memo attached to the entry. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupPrepaymentResponse := models.SubscriptionGroupPrepaymentResponse{
        Id:                   models.ToPointer(32),
        AmountInCents:        models.ToPointer(int64(138)),
        EndingBalanceInCents: models.ToPointer(int64(158)),
        EntryType:            models.ToPointer(models.ServiceCreditType_CREDIT),
        Memo:                 models.NewOptional(models.ToPointer("memo2")),
    }

}
```

