
# Subscription Group Response

## Structure

`SubscriptionGroupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroup`](../../doc/models/subscription-group.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupResponse := models.SubscriptionGroupResponse{
        SubscriptionGroup:    models.SubscriptionGroup{
            Uid:                     models.ToPointer("uid8"),
            CustomerId:              models.ToPointer(220),
            PaymentProfile:          models.ToPointer(models.SubscriptionGroupPaymentProfile{
                Id:                   models.ToPointer(44),
                FirstName:            models.ToPointer("first_name4"),
                LastName:             models.ToPointer("last_name2"),
                MaskedCardNumber:     models.ToPointer("masked_card_number2"),
            }),
            PaymentCollectionMethod: models.ToPointer(models.CollectionMethod_PREPAID),
            SubscriptionIds:         []int{
                74,
                75,
            },
        },
    }

}
```

