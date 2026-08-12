
# Subscription Migration Preview

## Structure

`SubscriptionMigrationPreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProratedAdjustmentInCents` | `*int64` | Optional | The amount of the prorated adjustment that would be issued for the current subscription. |
| `ChargeInCents` | `*int64` | Optional | The amount of the charge that would be created for the new product. |
| `PaymentDueInCents` | `*int64` | Optional | The amount of the payment due in the case of an upgrade. |
| `CreditAppliedInCents` | `*int64` | Optional | Represents a credit in cents that is applied to your subscription as part of a migration process for a specific product, which reduces the amount owed for the subscription. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMigrationPreview := models.SubscriptionMigrationPreview{
        ProratedAdjustmentInCents: models.ToPointer(int64(6)),
        ChargeInCents:             models.ToPointer(int64(144)),
        PaymentDueInCents:         models.ToPointer(int64(60)),
        CreditAppliedInCents:      models.ToPointer(int64(20)),
    }

}
```

