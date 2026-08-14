
# Create Invoice Address

Overrides the default address.

## Structure

`CreateInvoiceAddress`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Phone` | `*string` | Optional | - |
| `Address` | `*string` | Optional | - |
| `Address2` | `*string` | Optional | - |
| `City` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `Zip` | `*string` | Optional | - |
| `Country` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createInvoiceAddress := models.CreateInvoiceAddress{
        FirstName:            models.ToPointer("first_name6"),
        LastName:             models.ToPointer("last_name4"),
        Phone:                models.ToPointer("phone4"),
        Address:              models.ToPointer("address2"),
        Address2:             models.ToPointer("address_20"),
    }

}
```

