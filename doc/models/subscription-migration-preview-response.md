
# Subscription Migration Preview Response

## Structure

`SubscriptionMigrationPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Migration` | [`models.SubscriptionMigrationPreview`](../../doc/models/subscription-migration-preview.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMigrationPreviewResponse := models.SubscriptionMigrationPreviewResponse{
        Migration:            models.SubscriptionMigrationPreview{
            ProratedAdjustmentInCents: models.ToPointer(int64(196)),
            ChargeInCents:             models.ToPointer(int64(78)),
            PaymentDueInCents:         models.ToPointer(int64(250)),
            CreditAppliedInCents:      models.ToPointer(int64(210)),
        },
    }

}
```

