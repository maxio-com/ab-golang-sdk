
# Scheduled Renewal Item Request Body Component

## Structure

`ScheduledRenewalItemRequestBodyComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ItemType` | `string` | Required, Constant | Item type to add. Either Product or Component.<br><br>**Value**: `"Component"` |
| `ItemId` | `int` | Required | Product or component identifier. |
| `PricePointId` | `*int` | Optional | Price point identifier. |
| `Quantity` | `*int` | Optional | (Optional) Quantity for the item. |
| `CustomPrice` | [`*models.ScheduledRenewalComponentCustomPrice`](../../doc/models/scheduled-renewal-component-custom-price.md) | Optional | Custom pricing for a component within a scheduled renewal. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalItemRequestBodyComponent := models.ScheduledRenewalItemRequestBodyComponent{
        ItemType:             "Component",
        ItemId:               20,
        PricePointId:         models.ToPointer(6),
        Quantity:             models.ToPointer(84),
        CustomPrice:          models.ToPointer(models.ScheduledRenewalComponentCustomPrice{
            TaxIncluded:          models.ToPointer(false),
            PricingScheme:        models.PricingScheme_STAIRSTEP,
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
    }

}
```

