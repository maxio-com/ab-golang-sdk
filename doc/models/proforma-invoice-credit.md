
# Proforma Invoice Credit

## Structure

`ProformaInvoiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Memo` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `OriginalAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `AppliedAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proformaInvoiceCredit := models.ProformaInvoiceCredit{
        Uid:                  models.ToPointer("uid4"),
        Memo:                 models.ToPointer("memo8"),
        OriginalAmount:       models.ToPointer("original_amount8"),
        AppliedAmount:        models.ToPointer("applied_amount4"),
    }

}
```

