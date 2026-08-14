
# Create Metered Component

## Structure

`CreateMeteredComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MeteredComponent` | [`models.MeteredComponent`](../../doc/models/metered-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMeteredComponent := models.CreateMeteredComponent{
        MeteredComponent:     models.MeteredComponent{
            Name:                      "name0",
            UnitName:                  "unit_name2",
            Description:               models.ToPointer("description0"),
            Handle:                    models.ToPointer("handle6"),
            Taxable:                   models.ToPointer(false),
            PricingScheme:             models.PricingScheme_STAIRSTEP,
            Prices:                    []models.Price{
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
            PricePoints:               []models.ComponentPricePointItem{
                models.ComponentPricePointItem{
                    Name:                 models.ToPointer("name2"),
                    Handle:               models.ToPointer("handle8"),
                    PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
                    Interval:             models.ToPointer(92),
                    IntervalUnit:         models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                },
                models.ComponentPricePointItem{
                    Name:                 models.ToPointer("name2"),
                    Handle:               models.ToPointer("handle8"),
                    PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
                    Interval:             models.ToPointer(92),
                    IntervalUnit:         models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                },
                models.ComponentPricePointItem{
                    Name:                 models.ToPointer("name2"),
                    Handle:               models.ToPointer("handle8"),
                    PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
                    Interval:             models.ToPointer(92),
                    IntervalUnit:         models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
                },
            },
        },
    }

}
```

