
# Update Subscription Component

## Structure

`UpdateSubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `CustomPrice` | [`*models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionComponent := models.UpdateSubscriptionComponent{
        ComponentId:          models.ToPointer(118),
        CustomPrice:          models.ToPointer(models.ComponentCustomPrice{
            TaxIncluded:              models.ToPointer(false),
            PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
            Interval:                 models.ToPointer(66),
            IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
            ListPricePointId:         models.NewOptional(models.ToPointer(174)),
            Prices:                   []models.Price{
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
    }

}
```

