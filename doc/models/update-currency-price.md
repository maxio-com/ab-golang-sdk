
# Update Currency Price

## Structure

`UpdateCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | ID of the currency price record being updated |
| `Price` | `float64` | Required | New price for the given currency |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateCurrencyPrice := models.UpdateCurrencyPrice{
        Id:                   104,
        Price:                float64(163.6),
    }

}
```

