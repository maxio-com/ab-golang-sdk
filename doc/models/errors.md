
# Errors

## Structure

`Errors`

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
    mErrors := models.Errors{
        PerPage:              []string{
            "per_page9",
            "per_page8",
            "per_page7",
        },
        PricePoint:           []string{
            "price_point0",
            "price_point1",
        },
    }

}
```

