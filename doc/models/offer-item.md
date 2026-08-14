
# Offer Item

## Structure

`OfferItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `StartingQuantity` | `*string` | Optional | - |
| `Editable` | `*bool` | Optional | - |
| `ComponentUnitPrice` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `CurrencyPrices` | [`[]models.CurrencyPrice`](../../doc/models/currency-price.md) | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. e.g., an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    offerItem := models.OfferItem{
        ComponentId:          models.ToPointer(234),
        PricePointId:         models.ToPointer(254),
        StartingQuantity:     models.ToPointer("starting_quantity6"),
        Editable:             models.ToPointer(false),
        ComponentUnitPrice:   models.ToPointer("component_unit_price2"),
    }

}
```

