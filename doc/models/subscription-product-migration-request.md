
# Subscription Product Migration Request

## Structure

`SubscriptionProductMigrationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Migration` | [`models.SubscriptionProductMigration`](../../doc/models/subscription-product-migration.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionProductMigrationRequest := models.SubscriptionProductMigrationRequest{
        Migration:            models.SubscriptionProductMigration{
            ProductId:               models.ToPointer(158),
            ProductPricePointId:     models.ToPointer(82),
            IncludeTrial:            models.ToPointer(false),
            IncludeInitialCharge:    models.ToPointer(false),
            IncludeCoupons:          models.ToPointer(true),
            PreservePeriod:          models.ToPointer(false),
        },
    }

}
```

