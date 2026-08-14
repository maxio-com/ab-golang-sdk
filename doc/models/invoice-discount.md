
# Invoice Discount

## Structure

`InvoiceDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Code` | `*string` | Optional | - |
| `SourceType` | [`*models.InvoiceDiscountSourceType`](../../doc/models/invoice-discount-source-type.md) | Optional | - |
| `SourceId` | `*int` | Optional | - |
| `DiscountType` | [`*models.InvoiceDiscountType`](../../doc/models/invoice-discount-type.md) | Optional | - |
| `Percentage` | `*string` | Optional | - |
| `EligibleAmount` | `*string` | Optional | - |
| `DiscountAmount` | `*string` | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `LineItemBreakouts` | [`[]models.InvoiceDiscountBreakout`](../../doc/models/invoice-discount-breakout.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceDiscount := models.InvoiceDiscount{
        Uid:                  models.ToPointer("uid4"),
        Title:                models.ToPointer("title0"),
        Description:          models.NewOptional(models.ToPointer("description6")),
        Code:                 models.ToPointer("code2"),
        SourceType:           models.ToPointer(models.InvoiceDiscountSourceType_COUPON),
    }

}
```

