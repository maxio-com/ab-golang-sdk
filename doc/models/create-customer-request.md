
# Create Customer Request

## Structure

`CreateCustomerRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Customer` | [`models.CreateCustomer`](../../doc/models/create-customer.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createCustomerRequest := models.CreateCustomerRequest{
        Customer:             models.CreateCustomer{
            FirstName:            "first_name0",
            LastName:             "last_name8",
            Email:                "email6",
            CcEmails:             models.ToPointer("cc_emails0"),
            Organization:         models.ToPointer("organization6"),
            Reference:            models.ToPointer("reference4"),
            Address:              models.ToPointer("address6"),
            Address2:             models.ToPointer("address_24"),
        },
    }

}
```

