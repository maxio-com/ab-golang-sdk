
# Invoice Seller

Information about the seller (merchant) listed on the masthead of the invoice.

## Structure

`InvoiceSeller`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Address` | [`*models.InvoiceAddress`](../../doc/models/invoice-address.md) | Optional | - |
| `Phone` | `*string` | Optional | - |
| `LogoUrl` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceSeller := models.InvoiceSeller{
        Name:                 models.ToPointer("name4"),
        Address:              models.ToPointer(models.InvoiceAddress{
            Street:               models.NewOptional(models.ToPointer("street6")),
            Line2:                models.NewOptional(models.ToPointer("line20")),
            City:                 models.NewOptional(models.ToPointer("city6")),
            State:                models.NewOptional(models.ToPointer("state2")),
            Zip:                  models.NewOptional(models.ToPointer("zip0")),
        }),
        Phone:                models.ToPointer("phone6"),
        LogoUrl:              models.NewOptional(models.ToPointer("logo_url6")),
    }

}
```

