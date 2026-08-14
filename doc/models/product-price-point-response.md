
# Product Price Point Response

## Structure

`ProductPricePointResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.ProductPricePoint`](../../doc/models/product-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productPricePointResponse := models.ProductPricePointResponse{
        PricePoint:           models.ProductPricePoint{
            Id:                      models.ToPointer(248),
            Name:                    models.ToPointer("name0"),
            Handle:                  models.NewOptional(models.ToPointer("handle6")),
            PriceInCents:            models.ToPointer(int64(196)),
            Interval:                models.ToPointer(44),
        },
    }

}
```

