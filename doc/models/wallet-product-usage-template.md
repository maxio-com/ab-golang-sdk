
# Wallet Product Usage Template

One wallet template billing through this Wallet Product, joined to the Feature and plan/component that assigns it, when there is one.

## Structure

`WalletProductUsageTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PlanName` | `models.Optional[string]` | Optional | The name of the product or component the Feature is assigned to. `null` when the wallet template has no assigning Feature. |
| `WalletTemplateName` | `models.Optional[string]` | Optional | The label set in `wallet_behavior.name` on the Feature, or the standalone wallet's own `name`. |
| `UnitName` | `*string` | Optional | - |
| `PurchaseMechanism` | `models.Optional[string]` | Optional | A best-effort label for how this wallet is expected to be replenished. Currently reports `"Overage"` when the wallet template allows a negative balance, and `null` otherwise; there's no persisted field yet distinguishing, for example, an intended "Auto top-up" from "Purchase more" setup. Treat this as informational, not a stable enum. |
| `PricePointName` | `models.Optional[string]` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `Active` | `*bool` | Optional | `false` when the assigning Feature is archived. |
| `FeatureTemplateId` | `models.Optional[int]` | Optional | - |
| `FeatureTemplateName` | `models.Optional[string]` | Optional | - |
| `PricePointOverride` | `*bool` | Optional | `true` when this wallet template is priced at a price point other than the Wallet Product's default. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductUsageTemplate := models.WalletProductUsageTemplate{
        PlanName:             models.NewOptional(models.ToPointer("plan_name0")),
        WalletTemplateName:   models.NewOptional(models.ToPointer("wallet_template_name8")),
        UnitName:             models.ToPointer("unit_name0"),
        PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism8")),
        PricePointName:       models.NewOptional(models.ToPointer("price_point_name8")),
    }

}
```

