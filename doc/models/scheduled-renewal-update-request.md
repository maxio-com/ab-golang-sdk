
# Scheduled Renewal Update Request

## Structure

`ScheduledRenewalUpdateRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalConfigurationItem` | [`models.ScheduledRenewalUpdateRequestRenewalConfigurationItem`](../../doc/models/containers/scheduled-renewal-update-request-renewal-configuration-item.md) | Required | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalUpdateRequest := models.ScheduledRenewalUpdateRequest{
        RenewalConfigurationItem: models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
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

