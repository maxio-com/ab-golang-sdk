
# Customer Payer Change

## Structure

`CustomerPayerChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Before` | [`models.InvoicePayerChange`](../../doc/models/invoice-payer-change.md) | Required | - |
| `After` | [`models.InvoicePayerChange`](../../doc/models/invoice-payer-change.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerPayerChange := models.CustomerPayerChange{
        Before:               models.InvoicePayerChange{
            FirstName:            models.ToPointer("first_name0"),
            LastName:             models.ToPointer("last_name8"),
            Organization:         models.ToPointer("organization4"),
            Email:                models.ToPointer("email6"),
        },
        After:                models.InvoicePayerChange{
            FirstName:            models.ToPointer("first_name2"),
            LastName:             models.ToPointer("last_name0"),
            Organization:         models.ToPointer("organization4"),
            Email:                models.ToPointer("email4"),
        },
    }

}
```

