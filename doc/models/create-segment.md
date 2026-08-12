
# Create Segment

## Structure

`CreateSegment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SegmentProperty1Value` | [`*models.CreateSegmentSegmentProperty1Value`](../../doc/models/containers/create-segment-segment-property-1-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty2Value` | [`*models.CreateSegmentSegmentProperty2Value`](../../doc/models/containers/create-segment-segment-property-2-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty3Value` | [`*models.CreateSegmentSegmentProperty3Value`](../../doc/models/containers/create-segment-segment-property-3-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty4Value` | [`*models.CreateSegmentSegmentProperty4Value`](../../doc/models/containers/create-segment-segment-property-4-value.md) | Optional | This is a container for one-of cases. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSegment := models.CreateSegment{
        SegmentProperty1Value: models.ToPointer(models.CreateSegmentSegmentProperty1ValueContainer.FromString("String7")),
        SegmentProperty2Value: models.ToPointer(models.CreateSegmentSegmentProperty2ValueContainer.FromString("String9")),
        SegmentProperty3Value: models.ToPointer(models.CreateSegmentSegmentProperty3ValueContainer.FromString("String5")),
        SegmentProperty4Value: models.ToPointer(models.CreateSegmentSegmentProperty4ValueContainer.FromString("String1")),
        PricingScheme:         models.PricingScheme_STAIRSTEP,
        Prices:                []models.CreateOrUpdateSegmentPrice{
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
        },
    }

}
```

