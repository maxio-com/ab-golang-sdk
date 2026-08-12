
# Invoice Tax

## Structure

`InvoiceTax`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `SourceType` | [`*models.ProformaInvoiceTaxSourceType`](../../doc/models/proforma-invoice-tax-source-type.md) | Optional | - |
| `SourceId` | `*int` | Optional | - |
| `Percentage` | `*string` | Optional | - |
| `TaxableAmount` | `*string` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `LineItemBreakouts` | [`[]models.InvoiceTaxBreakout`](../../doc/models/invoice-tax-breakout.md) | Optional | - |
| `TaxComponentBreakouts` | [`[]models.InvoiceTaxComponentBreakout`](../../doc/models/invoice-tax-component-breakout.md) | Optional | - |
| `EuVat` | `*bool` | Optional | - |
| `Type` | `*string` | Optional | - |
| `TaxExemptAmount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceTax := models.InvoiceTax{
        Uid:                   models.ToPointer("uid2"),
        Title:                 models.ToPointer("title8"),
        Description:           models.NewOptional(models.ToPointer("description2")),
        SourceType:            models.ToPointer(models.ProformaInvoiceTaxSourceType_TAX),
        SourceId:              models.ToPointer(86),
    }

}
```

