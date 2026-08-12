
# Update Customer Request

## Structure

`UpdateCustomerRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Customer` | [`models.UpdateCustomer`](../../doc/models/update-customer.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateCustomerRequest := models.UpdateCustomerRequest{
        Customer:             models.UpdateCustomer{
            FirstName:            models.ToPointer("first_name0"),
            LastName:             models.ToPointer("last_name8"),
            Email:                models.ToPointer("email6"),
            CcEmails:             models.ToPointer("cc_emails0"),
            Organization:         models.ToPointer("organization6"),
        },
    }

}
```

