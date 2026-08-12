
# Scheduled Renewal Configuration Item Request

## Structure

`ScheduledRenewalConfigurationItemRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalConfigurationItem` | [`models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem`](../../doc/models/containers/scheduled-renewal-configuration-item-request-renewal-configuration-item.md) | Required | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalConfigurationItemRequest := models.ScheduledRenewalConfigurationItemRequest{
        RenewalConfigurationItem: models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
            ItemType:             "Component",
            ItemId:               108,
            PricePointId:         models.ToPointer(122),
            Quantity:             models.ToPointer(212),
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
        }),
    }

}
```

