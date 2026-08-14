
# Invoice Payer

## Structure

`InvoicePayer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyId` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `models.Optional[string]` | Optional | - |
| `Email` | `*string` | Optional | - |
| `VatNumber` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePayer := models.InvoicePayer{
        ChargifyId:           models.ToPointer(198),
        FirstName:            models.ToPointer("first_name2"),
        LastName:             models.ToPointer("last_name0"),
        Organization:         models.NewOptional(models.ToPointer("organization4")),
        Email:                models.ToPointer("email4"),
    }

}
```

