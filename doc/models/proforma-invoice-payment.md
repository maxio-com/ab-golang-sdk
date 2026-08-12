
# Proforma Invoice Payment

## Structure

`ProformaInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `OriginalAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `AppliedAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Prepayment` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proformaInvoicePayment := models.ProformaInvoicePayment{
        Memo:                 models.ToPointer("memo2"),
        OriginalAmount:       models.ToPointer("original_amount2"),
        AppliedAmount:        models.ToPointer("applied_amount0"),
        Prepayment:           models.ToPointer(false),
    }

}
```

