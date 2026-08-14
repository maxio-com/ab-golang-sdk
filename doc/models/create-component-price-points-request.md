
# Create Component Price Points Request

## Structure

`CreateComponentPricePointsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.CreateComponentPricePointsRequestPricePoints`](../../doc/models/containers/create-component-price-points-request-price-points.md) | Required | This is Array of a container for any-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createComponentPricePointsRequest := models.CreateComponentPricePointsRequest{
        PricePoints:          []models.CreateComponentPricePointsRequestPricePoints{
            models.CreateComponentPricePointsRequestPricePointsContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
                Name:                 "name0",
                Handle:               models.ToPointer("handle6"),
                PricingScheme:        models.PricingScheme_PERUNIT,
                Prices:               []models.Price{
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                },
                UseSiteExchangeRate:  models.ToPointer(false),
                TaxIncluded:          models.ToPointer(false),
                Interval:             models.ToPointer(24),
                IntervalUnit:         models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
            }),
        },
    }

}
```

