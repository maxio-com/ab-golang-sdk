
# Invoice Payment Method

## Structure

`InvoicePaymentMethod`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | `*string` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `CardBrand` | `*string` | Optional | - |
| `CardExpiration` | `*string` | Optional | - |
| `LastFour` | `models.Optional[string]` | Optional | - |
| `MaskedCardNumber` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePaymentMethod := models.InvoicePaymentMethod{
        Details:              models.ToPointer("details2"),
        Kind:                 models.ToPointer("kind0"),
        Memo:                 models.ToPointer("memo6"),
        Type:                 models.ToPointer("type8"),
        CardBrand:            models.ToPointer("card_brand4"),
    }

}
```

