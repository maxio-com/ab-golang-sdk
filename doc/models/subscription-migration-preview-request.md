
# Subscription Migration Preview Request

## Structure

`SubscriptionMigrationPreviewRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Migration` | [`models.SubscriptionMigrationPreviewOptions`](../../doc/models/subscription-migration-preview-options.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMigrationPreviewRequest := models.SubscriptionMigrationPreviewRequest{
        Migration:            models.SubscriptionMigrationPreviewOptions{
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

