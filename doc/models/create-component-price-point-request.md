
# Create Component Price Point Request

## Structure

`CreateComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CreateComponentPricePointRequestPricePoint`](../../doc/models/containers/create-component-price-point-request-price-point.md) | Required | This is a container for any-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createComponentPricePointRequest := models.CreateComponentPricePointRequest{
        PricePoint:           models.CreateComponentPricePointRequestPricePointContainer.FromCreateComponentPricePoint(models.CreateComponentPricePoint{
            Name:                   "name0",
            Handle:                 models.ToPointer("handle6"),
            PricingScheme:          models.PricingScheme_PERUNIT,
            Prices:                 []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
            },
            SecondaryPricing:       models.ToPointer(models.SecondaryPricing{
                PricingScheme:        models.ToPointer(models.PricingScheme_STAIRSTEP),
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
                },
            }),
            SecondaryPricingScheme: models.ToPointer(models.PricingScheme_STAIRSTEP),
            UseSiteExchangeRate:    models.ToPointer(true),
            TaxIncluded:            models.ToPointer(false),
        }),
    }

}
```

