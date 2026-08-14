
# Billing Manifest Item

## Structure

`BillingManifestItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | [`*models.LineItemTransactionType`](../../doc/models/line-item-transaction-type.md) | Optional | A handle for the line item transaction type |
| `Kind` | [`*models.BillingManifestLineItemKind`](../../doc/models/billing-manifest-line-item-kind.md) | Optional | A handle for the billing manifest line item kind |
| `AmountInCents` | `*int64` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `DiscountAmountInCents` | `*int64` | Optional | - |
| `TaxableAmountInCents` | `*int64` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductHandle` | `*string` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | - |
| `PeriodRangeEnd` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    billingManifestItem := models.BillingManifestItem{
        TransactionType:       models.ToPointer(models.LineItemTransactionType_PAYMENT),
        Kind:                  models.ToPointer(models.BillingManifestLineItemKind_TRIAL),
        AmountInCents:         models.ToPointer(int64(148)),
        Memo:                  models.ToPointer("memo0"),
        DiscountAmountInCents: models.ToPointer(int64(88)),
    }

}
```

