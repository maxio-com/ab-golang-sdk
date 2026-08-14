
# List Components Price Points Response

## Structure

`ListComponentsPricePointsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.ComponentPricePoint`](../../doc/models/component-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listComponentsPricePointsResponse := models.ListComponentsPricePointsResponse{
        PricePoints:          []models.ComponentPricePoint{
            models.ComponentPricePoint{
                Id:                       models.ToPointer(40),
                Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
                Default:                  models.ToPointer(false),
                Name:                     models.ToPointer("name2"),
                PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
            },
        },
    }

}
```

