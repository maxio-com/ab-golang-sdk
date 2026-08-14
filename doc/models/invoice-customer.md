
# Invoice Customer

Information about the customer who is owner or recipient of the invoiced subscription.

## Structure

`InvoiceCustomer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyId` | `models.Optional[int]` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `models.Optional[string]` | Optional | - |
| `Email` | `*string` | Optional | - |
| `VatNumber` | `models.Optional[string]` | Optional | - |
| `Reference` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceCustomer := models.InvoiceCustomer{
        ChargifyId:           models.NewOptional(models.ToPointer(82)),
        FirstName:            models.ToPointer("first_name2"),
        LastName:             models.ToPointer("last_name0"),
        Organization:         models.NewOptional(models.ToPointer("organization6")),
        Email:                models.ToPointer("email4"),
    }

}
```

