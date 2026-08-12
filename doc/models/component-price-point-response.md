
# Component Price Point Response

## Structure

`ComponentPricePointResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.ComponentPricePoint`](../../doc/models/component-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPricePointResponse := models.ComponentPricePointResponse{
        PricePoint:           models.ComponentPricePoint{
            Id:                       models.ToPointer(248),
            Type:                     models.ToPointer(models.PricePointType_ENUMDEFAULT),
            Default:                  models.ToPointer(false),
            Name:                     models.ToPointer("name0"),
            PricingScheme:            models.ToPointer(models.PricingScheme_PERUNIT),
        },
    }

}
```

