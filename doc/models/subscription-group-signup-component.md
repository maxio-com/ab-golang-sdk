
# Subscription Group Signup Component

## Structure

`SubscriptionGroupSignupComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | [`*models.SubscriptionGroupSignupComponentComponentId`](../../doc/models/containers/subscription-group-signup-component-component-id.md) | Optional | This is a container for one-of cases. |
| `AllocatedQuantity` | [`*models.SubscriptionGroupSignupComponentAllocatedQuantity`](../../doc/models/containers/subscription-group-signup-component-allocated-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitBalance` | [`*models.SubscriptionGroupSignupComponentUnitBalance`](../../doc/models/containers/subscription-group-signup-component-unit-balance.md) | Optional | This is a container for one-of cases. |
| `PricePointId` | [`*models.SubscriptionGroupSignupComponentPricePointId`](../../doc/models/containers/subscription-group-signup-component-price-point-id.md) | Optional | This is a container for one-of cases. |
| `CustomPrice` | [`*models.SubscriptionGroupComponentCustomPrice`](../../doc/models/subscription-group-component-custom-price.md) | Optional | Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupSignupComponent := models.SubscriptionGroupSignupComponent{
        ComponentId:          models.ToPointer(models.SubscriptionGroupSignupComponentComponentIdContainer.FromString("String7")),
        AllocatedQuantity:    models.ToPointer(models.SubscriptionGroupSignupComponentAllocatedQuantityContainer.FromString("String1")),
        UnitBalance:          models.ToPointer(models.SubscriptionGroupSignupComponentUnitBalanceContainer.FromString("String5")),
        PricePointId:         models.ToPointer(models.SubscriptionGroupSignupComponentPricePointIdContainer.FromString("String1")),
        CustomPrice:          models.ToPointer(models.SubscriptionGroupComponentCustomPrice{
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
        }),
    }

}
```

