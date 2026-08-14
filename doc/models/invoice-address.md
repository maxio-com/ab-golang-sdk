
# Invoice Address

## Structure

`InvoiceAddress`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Street` | `models.Optional[string]` | Optional | - |
| `Line2` | `models.Optional[string]` | Optional | - |
| `City` | `models.Optional[string]` | Optional | - |
| `State` | `models.Optional[string]` | Optional | - |
| `Zip` | `models.Optional[string]` | Optional | - |
| `Country` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceAddress := models.InvoiceAddress{
        Street:               models.NewOptional(models.ToPointer("street2")),
        Line2:                models.NewOptional(models.ToPointer("line26")),
        City:                 models.NewOptional(models.ToPointer("city2")),
        State:                models.NewOptional(models.ToPointer("state8")),
        Zip:                  models.NewOptional(models.ToPointer("zip6")),
    }

}
```

