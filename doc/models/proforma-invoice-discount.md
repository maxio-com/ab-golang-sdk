
# Proforma Invoice Discount

## Structure

`ProformaInvoiceDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Code` | `*string` | Optional | - |
| `SourceType` | [`*models.ProformaInvoiceDiscountSourceType`](../../doc/models/proforma-invoice-discount-source-type.md) | Optional | - |
| `DiscountType` | [`*models.InvoiceDiscountType`](../../doc/models/invoice-discount-type.md) | Optional | - |
| `EligibleAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DiscountAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.InvoiceDiscountBreakout`](../../doc/models/invoice-discount-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proformaInvoiceDiscount := models.ProformaInvoiceDiscount{
        Uid:                  models.ToPointer("uid2"),
        Title:                models.ToPointer("title8"),
        Code:                 models.ToPointer("code0"),
        SourceType:           models.ToPointer(models.ProformaInvoiceDiscountSourceType_COUPON),
        DiscountType:         models.ToPointer(models.InvoiceDiscountType_PERCENTAGE),
    }

}
```

