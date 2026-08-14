
# Billing Manifest

## Structure

`BillingManifest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LineItems` | [`[]models.BillingManifestItem`](../../doc/models/billing-manifest-item.md) | Optional | - |
| `TotalInCents` | `*int64` | Optional | - |
| `TotalDiscountInCents` | `*int64` | Optional | - |
| `TotalTaxInCents` | `*int64` | Optional | - |
| `SubtotalInCents` | `*int64` | Optional | - |
| `StartDate` | `models.Optional[time.Time]` | Optional | - |
| `EndDate` | `models.Optional[time.Time]` | Optional | - |
| `PeriodType` | `models.Optional[string]` | Optional | - |
| `ExistingBalanceInCents` | `*int64` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    billingManifest := models.BillingManifest{
        LineItems:              []models.BillingManifestItem{
            models.BillingManifestItem{
                TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
                Kind:                  models.ToPointer(models.BillingManifestLineItemKind_COMPONENT),
                AmountInCents:         models.ToPointer(int64(24)),
                Memo:                  models.ToPointer("memo2"),
                DiscountAmountInCents: models.ToPointer(int64(172)),
            },
        },
        TotalInCents:           models.ToPointer(int64(96)),
        TotalDiscountInCents:   models.ToPointer(int64(174)),
        TotalTaxInCents:        models.ToPointer(int64(76)),
        SubtotalInCents:        models.ToPointer(int64(208)),
    }

}
```

