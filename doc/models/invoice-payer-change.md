
# Invoice Payer Change

## Structure

`InvoicePayerChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `*string` | Optional | - |
| `Email` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePayerChange := models.InvoicePayerChange{
        FirstName:            models.ToPointer("first_name4"),
        LastName:             models.ToPointer("last_name2"),
        Organization:         models.ToPointer("organization2"),
        Email:                models.ToPointer("email2"),
    }

}
```

