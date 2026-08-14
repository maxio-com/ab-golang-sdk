
# Invoice Tax Component Breakout

## Structure

`InvoiceTaxComponentBreakout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TaxRuleId` | `*int` | Optional | - |
| `Percentage` | `*string` | Optional | - |
| `CountryCode` | `*string` | Optional | - |
| `SubdivisionCode` | `*string` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `TaxableAmount` | `*string` | Optional | - |
| `TaxExemptAmount` | `*string` | Optional | - |
| `NonTaxableAmount` | `*string` | Optional | - |
| `TaxName` | `*string` | Optional | - |
| `TaxType` | `*string` | Optional | - |
| `RateType` | `*string` | Optional | - |
| `TaxAuthorityType` | `*int` | Optional | - |
| `StateAssignedNo` | `*string` | Optional | - |
| `TaxSubType` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceTaxComponentBreakout := models.InvoiceTaxComponentBreakout{
        TaxRuleId:            models.ToPointer(66),
        Percentage:           models.ToPointer("percentage0"),
        CountryCode:          models.ToPointer("country_code2"),
        SubdivisionCode:      models.ToPointer("subdivision_code6"),
        TaxAmount:            models.ToPointer("tax_amount4"),
    }

}
```

