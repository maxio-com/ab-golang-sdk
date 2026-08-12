
# Subscription Group Component Custom Price

Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`.

## Structure

`SubscriptionGroupComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | - |
| `OveragePricing` | [`[]models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupComponentCustomPrice := models.SubscriptionGroupComponentCustomPrice{
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
        OveragePricing:       []models.ComponentCustomPrice{
            models.ComponentCustomPrice{
                TaxIncluded:              models.ToPointer(false),
                PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
                Interval:                 models.ToPointer(230),
                IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                ListPricePointId:         models.NewOptional(models.ToPointer(10)),
                Prices:                   []models.Price{
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                },
            },
            models.ComponentCustomPrice{
                TaxIncluded:              models.ToPointer(false),
                PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
                Interval:                 models.ToPointer(230),
                IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                ListPricePointId:         models.NewOptional(models.ToPointer(10)),
                Prices:                   []models.Price{
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                },
            },
            models.ComponentCustomPrice{
                TaxIncluded:              models.ToPointer(false),
                PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
                Interval:                 models.ToPointer(230),
                IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                ListPricePointId:         models.NewOptional(models.ToPointer(10)),
                Prices:                   []models.Price{
                    models.Price{
                        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                    },
                },
            },
        },
    }

}
```

