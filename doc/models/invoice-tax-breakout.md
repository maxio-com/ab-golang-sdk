
# Invoice Tax Breakout

## Structure

`InvoiceTaxBreakout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `TaxableAmount` | `*string` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `TaxExemptAmount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceTaxBreakout := models.InvoiceTaxBreakout{
        Uid:                  models.ToPointer("uid4"),
        TaxableAmount:        models.ToPointer("taxable_amount8"),
        TaxAmount:            models.ToPointer("tax_amount2"),
        TaxExemptAmount:      models.ToPointer("tax_exempt_amount4"),
    }

}
```

