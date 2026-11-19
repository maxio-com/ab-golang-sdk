
# Errors

## Structure

`Errors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | `[]string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    mErrors := models.Errors{
        PricePoint:           []string{
            "price_point0",
            "price_point1",
        },
    }

}
```

