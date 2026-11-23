
# Segment

## Structure

`Segment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `EventBasedBillingMetricId` | `*int` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `SegmentProperty1Value` | [`*models.SegmentSegmentProperty1Value`](../../doc/models/containers/segment-segment-property-1-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty2Value` | [`*models.SegmentSegmentProperty2Value`](../../doc/models/containers/segment-segment-property-2-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty3Value` | [`*models.SegmentSegmentProperty3Value`](../../doc/models/containers/segment-segment-property-3-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty4Value` | [`*models.SegmentSegmentProperty4Value`](../../doc/models/containers/segment-segment-property-4-value.md) | Optional | This is a container for one-of cases. |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `Prices` | [`[]models.SegmentPrice`](../../doc/models/segment-price.md) | Optional | **Constraints**: *Minimum Items*: `1` |
| `CurrencyPrices` | [`[]models.ComponentCurrencyPrice`](../../doc/models/component-currency-price.md) | Optional | Prices of the segment in the site's non-primary currencies, one entry for each price and currency. This field is present only when the request includes `currency_prices=true`. If the price point uses the site exchange rate, the prices are calculated from the current rate and each `id` is `null`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    segment := models.Segment{
        Id:                        models.ToPointer(118),
        ComponentId:               models.ToPointer(228),
        PricePointId:              models.ToPointer(4),
        EventBasedBillingMetricId: models.ToPointer(56),
        PricingScheme:             models.ToPointer(models.PricingScheme_STAIRSTEP),
    }

}
```

