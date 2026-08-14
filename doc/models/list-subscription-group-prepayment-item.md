
# List Subscription Group Prepayment Item

## Structure

`ListSubscriptionGroupPrepaymentItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SubscriptionGroupUid` | `*string` | Optional | - |
| `AmountInCents` | `*int64` | Optional | - |
| `RemainingAmountInCents` | `*int64` | Optional | - |
| `Details` | `*string` | Optional | - |
| `External` | `*bool` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `PaymentType` | [`*models.PrepaymentMethod`](../../doc/models/prepayment-method.md) | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupPrepaymentItem := models.ListSubscriptionGroupPrepaymentItem{
        Id:                     models.ToPointer(72),
        SubscriptionGroupUid:   models.ToPointer("subscription_group_uid6"),
        AmountInCents:          models.ToPointer(int64(98)),
        RemainingAmountInCents: models.ToPointer(int64(216)),
        Details:                models.ToPointer("details2"),
    }

}
```

