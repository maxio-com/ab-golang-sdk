
# Proforma Invoice Tax

## Structure

`ProformaInvoiceTax`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Title` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SourceType` | [`*models.ProformaInvoiceTaxSourceType`](../../doc/models/proforma-invoice-tax-source-type.md) | Optional | The source of the tax. `Tax` is a tax configured in Advanced Billing. The other values identify the tax service that calculated the tax. |
| `Percentage` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxableAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItemBreakouts` | [`[]models.InvoiceTaxBreakout`](../../doc/models/invoice-tax-breakout.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proformaInvoiceTax := models.ProformaInvoiceTax{
        Uid:                  models.ToPointer("uid4"),
        Title:                models.ToPointer("title0"),
        SourceType:           models.ToPointer(models.ProformaInvoiceTaxSourceType_NUMERAL),
        Percentage:           models.ToPointer("percentage2"),
        TaxableAmount:        models.ToPointer("taxable_amount8"),
    }

}
```

