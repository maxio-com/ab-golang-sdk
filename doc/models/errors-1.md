
# Errors 1

## Structure

`Errors1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PerPage` | `[]string` | Optional | - |
| `PricePoint` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    errors1 := models.Errors1{
        PerPage:              []string{
            "per_page9",
        },
        PricePoint:           []string{
            "price_point0",
        },
    }

}
```

