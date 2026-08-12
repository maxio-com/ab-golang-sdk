
# Scheduled Renewal Component Custom Price

Custom pricing for a component within a scheduled renewal.

## Structure

`ScheduledRenewalComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TaxIncluded` | `*bool` | Optional | Whether or not the price point includes tax |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | Omit for On/Off components. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Required | On/off components only need one price bracket starting at 1. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalComponentCustomPrice := models.ScheduledRenewalComponentCustomPrice{
        TaxIncluded:          models.ToPointer(false),
        PricingScheme:        models.PricingScheme_PERUNIT,
        Prices:               []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
        },
    }

}
```

