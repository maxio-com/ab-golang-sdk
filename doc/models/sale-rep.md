
# Sale Rep

## Structure

`SaleRep`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FullName` | `*string` | Optional | - |
| `SubscriptionsCount` | `*int` | Optional | - |
| `TestMode` | `*bool` | Optional | - |
| `Subscriptions` | [`[]models.SaleRepSubscription`](../../doc/models/sale-rep-subscription.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    saleRep := models.SaleRep{
        Id:                   models.ToPointer(78),
        FullName:             models.ToPointer("full_name0"),
        SubscriptionsCount:   models.ToPointer(102),
        TestMode:             models.ToPointer(false),
        Subscriptions:        []models.SaleRepSubscription{
            models.SaleRepSubscription{
                Id:                   models.ToPointer(202),
                SiteName:             models.ToPointer("site_name8"),
                SubscriptionUrl:      models.ToPointer("subscription_url2"),
                CustomerName:         models.ToPointer("customer_name8"),
                CreatedAt:            models.ToPointer("created_at4"),
            },
            models.SaleRepSubscription{
                Id:                   models.ToPointer(202),
                SiteName:             models.ToPointer("site_name8"),
                SubscriptionUrl:      models.ToPointer("subscription_url2"),
                CustomerName:         models.ToPointer("customer_name8"),
                CreatedAt:            models.ToPointer("created_at4"),
            },
            models.SaleRepSubscription{
                Id:                   models.ToPointer(202),
                SiteName:             models.ToPointer("site_name8"),
                SubscriptionUrl:      models.ToPointer("subscription_url2"),
                CustomerName:         models.ToPointer("customer_name8"),
                CreatedAt:            models.ToPointer("created_at4"),
            },
        },
    }

}
```

