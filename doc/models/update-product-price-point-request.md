
# Update Product Price Point Request

## Structure

`UpdateProductPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.UpdateProductPricePoint`](../../doc/models/update-product-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateProductPricePointRequest := models.UpdateProductPricePointRequest{
        PricePoint:           models.UpdateProductPricePoint{
            Handle:               models.ToPointer("handle6"),
            PriceInCents:         models.ToPointer(int64(196)),
        },
    }

}
```

