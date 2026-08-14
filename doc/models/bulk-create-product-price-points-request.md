
# Bulk Create Product Price Points Request

## Structure

`BulkCreateProductPricePointsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.CreateProductPricePoint`](../../doc/models/create-product-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bulkCreateProductPricePointsRequest := models.BulkCreateProductPricePointsRequest{
        PricePoints:          []models.CreateProductPricePoint{
            models.CreateProductPricePoint{
                Name:                    "name2",
                Handle:                  models.ToPointer("handle8"),
                PriceInCents:            int64(108),
                Interval:                92,
                IntervalUnit:            models.IntervalUnit_DAY,
                TrialPriceInCents:       models.ToPointer(int64(196)),
                TrialInterval:           models.ToPointer(250),
                TrialIntervalUnit:       models.ToPointer(models.IntervalUnit_DAY),
                TrialType:               models.NewOptional(models.ToPointer(models.TrialType_NOOBLIGATION)),
                UseSiteExchangeRate:     models.ToPointer(true),
            },
        },
    }

}
```

