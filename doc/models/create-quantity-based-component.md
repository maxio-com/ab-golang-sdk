
# Create Quantity Based Component

## Structure

`CreateQuantityBasedComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `QuantityBasedComponent` | [`models.QuantityBasedComponent`](../../doc/models/quantity-based-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createQuantityBasedComponent := models.CreateQuantityBasedComponent{
        QuantityBasedComponent: models.QuantityBasedComponent{
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
            },
            UpgradeCharge:             models.NewOptional(models.ToPointer(models.CreditType_PRORATED)),
        },
    }

}
```

