
# Component Custom Price

Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.

## Structure

`ComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TaxIncluded` | `*bool` | Optional | Whether or not the price point includes tax |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | Omit for On/Off components. |
| `Interval` | `*int` | Optional | The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `ListPricePointId` | `models.Optional[int]` | Optional | (Optional) Id of the price point to use for list price calculations when<br>overriding the customer price. |
| `UseDefaultListPrice` | `*bool` | Optional | When true, list price calculations will continue to use the default price point even when a `custom_price` is supplied. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Required | On/off components only need one price bracket starting at 1. |
| `RenewPrepaidAllocation` | `*bool` | Optional | Applicable only to prepaid usage components. Controls whether the allocated quantity renews each period. |
| `RolloverPrepaidRemainder` | `*bool` | Optional | Applicable only to prepaid usage components. Controls whether remaining units roll over to the next period. |
| `ExpirationInterval` | `models.Optional[int]` | Optional | Applicable only when rollover is enabled. Number of `expiration_interval_unit`s after which rollover amounts expire. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | Applicable only when rollover is enabled. Interval unit for rollover expiration (month or day). |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentCustomPrice := models.ComponentCustomPrice{
        TaxIncluded:              models.ToPointer(false),
        PricingScheme:            models.ToPointer(models.PricingScheme_STAIRSTEP),
        Interval:                 models.ToPointer(58),
        IntervalUnit:             models.NewOptional(models.ToPointer(models.IntervalUnit_DAY)),
        ListPricePointId:         models.NewOptional(models.ToPointer(182)),
        Prices:                   []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
        },
    }

}
```

