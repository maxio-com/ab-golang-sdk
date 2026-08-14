
# Allocation Preview Line Item

## Structure

`AllocationPreviewLineItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionType` | [`*models.LineItemTransactionType`](../../doc/models/line-item-transaction-type.md) | Optional | A handle for the line item transaction type |
| `Kind` | [`*models.AllocationPreviewLineItemKind`](../../doc/models/allocation-preview-line-item-kind.md) | Optional | A handle for the line item kind for allocation preview |
| `AmountInCents` | `*int64` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `DiscountAmountInCents` | `*int64` | Optional | - |
| `TaxableAmountInCents` | `*int64` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `ComponentHandle` | `*string` | Optional | - |
| `Direction` | [`*models.AllocationPreviewDirection`](../../doc/models/allocation-preview-direction.md) | Optional | Visible when using Fine-grained Component Control. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    allocationPreviewLineItem := models.AllocationPreviewLineItem{
        TransactionType:       models.ToPointer(models.LineItemTransactionType_CREDIT),
        Kind:                  models.ToPointer(models.AllocationPreviewLineItemKind_QUANTITYBASEDCOMPONENT),
        AmountInCents:         models.ToPointer(int64(236)),
        Memo:                  models.ToPointer("memo6"),
        DiscountAmountInCents: models.ToPointer(int64(40)),
    }

}
```

