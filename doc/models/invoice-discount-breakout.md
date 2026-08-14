
# Invoice Discount Breakout

## Structure

`InvoiceDiscountBreakout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `EligibleAmount` | `*string` | Optional | - |
| `DiscountAmount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceDiscountBreakout := models.InvoiceDiscountBreakout{
        Uid:                  models.ToPointer("uid0"),
        EligibleAmount:       models.ToPointer("eligible_amount2"),
        DiscountAmount:       models.ToPointer("discount_amount4"),
    }

}
```

