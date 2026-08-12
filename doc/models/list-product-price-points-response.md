
# List Product Price Points Response

## Structure

`ListProductPricePointsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.ProductPricePoint`](../../doc/models/product-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listProductPricePointsResponse := models.ListProductPricePointsResponse{
        PricePoints:          []models.ProductPricePoint{
            models.ProductPricePoint{
                Id:                      models.ToPointer(40),
                Name:                    models.ToPointer("name2"),
                Handle:                  models.NewOptional(models.ToPointer("handle8")),
                PriceInCents:            models.ToPointer(int64(108)),
                Interval:                models.ToPointer(92),
            },
        },
    }

}
```

