
# Renewal Preview Line Item

## Structure

`RenewalPreviewLineItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | [`*models.LineItemTransactionType`](../../doc/models/line-item-transaction-type.md) | Optional | A handle for the line item transaction type |
| `Kind` | [`*models.LineItemKind`](../../doc/models/line-item-kind.md) | Optional | A handle for the line item kind |
| `AmountInCents` | `*int64` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `DiscountAmountInCents` | `*int64` | Optional | - |
| `TaxableAmountInCents` | `*int64` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `ProductHandle` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | - |
| `PeriodRangeEnd` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    renewalPreviewLineItem := models.RenewalPreviewLineItem{
        TransactionType:       models.ToPointer(models.LineItemTransactionType_ADJUSTMENT),
        Kind:                  models.ToPointer(models.LineItemKind_PREPAIDUSAGECOMPONENT),
        AmountInCents:         models.ToPointer(int64(32)),
        Memo:                  models.ToPointer("memo0"),
        DiscountAmountInCents: models.ToPointer(int64(228)),
    }

}
```

