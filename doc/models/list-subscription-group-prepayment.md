
# List Subscription Group Prepayment

## Structure

`ListSubscriptionGroupPrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.ListSubscriptionGroupPrepaymentItem`](../../doc/models/list-subscription-group-prepayment-item.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupPrepayment := models.ListSubscriptionGroupPrepayment{
        Prepayment:           models.ListSubscriptionGroupPrepaymentItem{
            Id:                     models.ToPointer(38),
            SubscriptionGroupUid:   models.ToPointer("subscription_group_uid2"),
            AmountInCents:          models.ToPointer(int64(124)),
            RemainingAmountInCents: models.ToPointer(int64(182)),
            Details:                models.ToPointer("details8"),
        },
    }

}
```

