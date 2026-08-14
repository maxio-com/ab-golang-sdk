
# Subscription Response

## Structure

`SubscriptionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`*models.Subscription`](../../doc/models/subscription.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionResponse := models.SubscriptionResponse{
        Subscription:         models.ToPointer(models.Subscription{
            Id:                                models.ToPointer(8),
            State:                             models.ToPointer(models.SubscriptionState_PAUSED),
            BalanceInCents:                    models.ToPointer(int64(124)),
            TotalRevenueInCents:               models.ToPointer(int64(48)),
            ProductPriceInCents:               models.ToPointer(int64(238)),
        }),
    }

}
```

