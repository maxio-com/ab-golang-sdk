
# Update Component Price Point Request

## Structure

`UpdateComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`*models.UpdateComponentPricePoint`](../../doc/models/update-component-price-point.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateComponentPricePointRequest := models.UpdateComponentPricePointRequest{
        PricePoint:           models.ToPointer(models.UpdateComponentPricePoint{
            Name:                 models.ToPointer("name0"),
            Handle:               models.ToPointer("handle6"),
            PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
            UseSiteExchangeRate:  models.ToPointer(false),
            TaxIncluded:          models.ToPointer(false),
        }),
    }

}
```

