
# Address Change

## Structure

`AddressChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Before` | [`models.InvoiceAddress`](../../doc/models/invoice-address.md) | Required | - |
| `After` | [`models.InvoiceAddress`](../../doc/models/invoice-address.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    addressChange := models.AddressChange{
        Before:               models.InvoiceAddress{
            Street:               models.NewOptional(models.ToPointer("street0")),
            Line2:                models.NewOptional(models.ToPointer("line24")),
            City:                 models.NewOptional(models.ToPointer("city0")),
            State:                models.NewOptional(models.ToPointer("state6")),
            Zip:                  models.NewOptional(models.ToPointer("zip4")),
        },
        After:                models.InvoiceAddress{
            Street:               models.NewOptional(models.ToPointer("street2")),
            Line2:                models.NewOptional(models.ToPointer("line26")),
            City:                 models.NewOptional(models.ToPointer("city8")),
            State:                models.NewOptional(models.ToPointer("state2")),
            Zip:                  models.NewOptional(models.ToPointer("zip4")),
        },
    }

}
```

