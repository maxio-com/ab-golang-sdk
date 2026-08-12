
# Component Price Points Response

## Structure

`ComponentPricePointsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.ComponentPricePoint`](../../doc/models/component-price-point.md) | Optional | - |
| `Meta` | [`*models.ListPublicKeysMeta`](../../doc/models/list-public-keys-meta.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPricePointsResponse := models.ComponentPricePointsResponse{
        PricePoints:          []models.ComponentPricePoint{
            models.ComponentPricePoint{
                Id:                       models.ToPointer(40),
                Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
                Default:                  models.ToPointer(false),
                Name:                     models.ToPointer("name2"),
                PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
            },
            models.ComponentPricePoint{
                Id:                       models.ToPointer(40),
                Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
                Default:                  models.ToPointer(false),
                Name:                     models.ToPointer("name2"),
                PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
            },
            models.ComponentPricePoint{
                Id:                       models.ToPointer(40),
                Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
                Default:                  models.ToPointer(false),
                Name:                     models.ToPointer("name2"),
                PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
            },
        },
        Meta:                 models.ToPointer(models.ListPublicKeysMeta{
            TotalCount:           models.ToPointer(150),
            CurrentPage:          models.ToPointer(126),
            TotalPages:           models.ToPointer(138),
            PerPage:              models.ToPointer(152),
        }),
    }

}
```

