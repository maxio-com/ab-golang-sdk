
# Create Prepaid Component

## Structure

`CreatePrepaidComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidUsageComponent` | [`models.PrepaidUsageComponent`](../../doc/models/prepaid-usage-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPrepaidComponent := models.CreatePrepaidComponent{
        PrepaidUsageComponent: models.PrepaidUsageComponent{
            Name:                      "name2",
            UnitName:                  "unit_name4",
            Description:               models.ToPointer("description2"),
            Handle:                    models.ToPointer("handle8"),
            Taxable:                   models.ToPointer(false),
            PricingScheme:             models.PricingScheme_PERUNIT,
            Prices:                    []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
            },
            UpgradeCharge:             models.NewOptional(models.ToPointer(models.CreditType_FULL)),
            OveragePricing:            models.OveragePricing{
                PricingScheme:        models.PricingScheme_STAIRSTEP,
                Prices:               []models.Price{
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

