
# Create Product Price Point Request

## Structure

`CreateProductPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CreateProductPricePoint`](../../doc/models/create-product-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createProductPricePointRequest := models.CreateProductPricePointRequest{
        PricePoint:           models.CreateProductPricePoint{
            Name:                    "name0",
            Handle:                  models.ToPointer("handle6"),
            PriceInCents:            int64(196),
            Interval:                44,
            IntervalUnit:            models.IntervalUnit_DAY,
            TrialPriceInCents:       models.ToPointer(int64(108)),
            TrialInterval:           models.ToPointer(202),
            TrialIntervalUnit:       models.ToPointer(models.IntervalUnit_DAY),
            TrialType:               models.NewOptional(models.ToPointer(models.TrialType_NOOBLIGATION)),
            UseSiteExchangeRate:     models.ToPointer(true),
        },
    }

}
```

