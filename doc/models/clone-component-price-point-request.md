
# Clone Component Price Point Request

## Structure

`CloneComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CloneComponentPricePoint`](../../doc/models/clone-component-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cloneComponentPricePointRequest := models.CloneComponentPricePointRequest{
        PricePoint:           models.CloneComponentPricePoint{
            Name:                 "name0",
            Handle:               models.ToPointer("handle6"),
        },
    }

}
```

