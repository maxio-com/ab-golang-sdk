
# Subscription Component Custom Price Segment

Custom rates for one segment of an event-based component, set on the subscription through `custom_price.segments`.

Identify the segment by the property values that occur in your events, in the same positions as the segmenting properties defined on the component's meter. Each segment in a request must have a unique combination of property values.

Property values are matched against your events as text. Send values for boolean segmenting properties as the strings `"true"` and `"false"`: JSON booleans are not supported here, and a segment whose only property value is `false` is rejected with a 422.

## Structure

`SubscriptionComponentCustomPriceSegment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SegmentProperty1Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty1Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-1-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty2Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty2Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-2-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty3Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty3Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-3-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty4Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty4Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-4-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty5Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty5Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-5-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty6Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty6Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-6-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty7Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty7Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-7-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty8Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty8Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-8-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty9Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty9Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-9-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty10Value` | [`*models.SubscriptionComponentCustomPriceSegmentSegmentProperty10Value`](../../doc/models/containers/subscription-component-custom-price-segment-segment-property-10-value.md) | Optional | This is a container for one-of cases. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Required | The rates for this segment. Use one price for a flat rate, or several price brackets for a tiered, volume or stairstep scheme.<br><br>**Constraints**: *Minimum Items*: `1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionComponentCustomPriceSegment := models.SubscriptionComponentCustomPriceSegment{
        SegmentProperty1Value:  models.ToPointer(models.SubscriptionComponentCustomPriceSegmentSegmentProperty1ValueContainer.FromString("String3")),
        SegmentProperty2Value:  models.ToPointer(models.SubscriptionComponentCustomPriceSegmentSegmentProperty2ValueContainer.FromString("String5")),
        SegmentProperty3Value:  models.ToPointer(models.SubscriptionComponentCustomPriceSegmentSegmentProperty3ValueContainer.FromString("String3")),
        SegmentProperty4Value:  models.ToPointer(models.SubscriptionComponentCustomPriceSegmentSegmentProperty4ValueContainer.FromString("String7")),
        SegmentProperty5Value:  models.ToPointer(models.SubscriptionComponentCustomPriceSegmentSegmentProperty5ValueContainer.FromString("String7")),
        PricingScheme:          models.PricingScheme_PERUNIT,
        Prices:                 []models.CreateOrUpdateSegmentPrice{
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
        },
    }

}
```

