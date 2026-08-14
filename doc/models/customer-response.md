
# Customer Response

## Structure

`CustomerResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Customer` | [`models.Customer`](../../doc/models/customer.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerResponse := models.CustomerResponse{
        Customer:             models.Customer{
            FirstName:                   models.ToPointer("first_name0"),
            LastName:                    models.ToPointer("last_name8"),
            Email:                       models.ToPointer("email6"),
            CcEmails:                    models.NewOptional(models.ToPointer("cc_emails0")),
            Organization:                models.NewOptional(models.ToPointer("organization6")),
        },
    }

}
```

