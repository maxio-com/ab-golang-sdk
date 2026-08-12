
# Invoice Balance Item

## Structure

`InvoiceBalanceItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Number` | `*string` | Optional | - |
| `OutstandingAmount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceBalanceItem := models.InvoiceBalanceItem{
        Uid:                  models.ToPointer("uid8"),
        Number:               models.ToPointer("number6"),
        OutstandingAmount:    models.ToPointer("outstanding_amount6"),
    }

}
```

