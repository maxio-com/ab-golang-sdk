
# List Subscription Group Prepayment Response

## Structure

`ListSubscriptionGroupPrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`[]models.ListSubscriptionGroupPrepayment`](../../doc/models/list-subscription-group-prepayment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupPrepaymentResponse := models.ListSubscriptionGroupPrepaymentResponse{
        Prepayments:          []models.ListSubscriptionGroupPrepayment{
            models.ListSubscriptionGroupPrepayment{
                Prepayment:           models.ListSubscriptionGroupPrepaymentItem{
                    Id:                     models.ToPointer(38),
                    SubscriptionGroupUid:   models.ToPointer("subscription_group_uid2"),
                    AmountInCents:          models.ToPointer(int64(124)),
                    RemainingAmountInCents: models.ToPointer(int64(182)),
                    Details:                models.ToPointer("details8"),
                },
            },
        },
    }

}
```

