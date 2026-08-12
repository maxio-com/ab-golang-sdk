
# Create EBB Component

## Structure

`CreateEBBComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EventBasedComponent` | [`models.EBBComponent`](../../doc/models/ebb-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createEBBComponent := models.CreateEBBComponent{
        EventBasedComponent:  models.EBBComponent{
            Name:                      "name8",
            UnitName:                  "unit_name0",
            Description:               models.ToPointer("description8"),
            Handle:                    models.ToPointer("handle4"),
            Taxable:                   models.ToPointer(false),
            PricingScheme:             models.PricingScheme_STAIRSTEP,
            Prices:                    []models.Price{
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
            },
            EventBasedBillingMetricId: 68,
        },
    }

}
```

