
# Create Segment Request

## Structure

`CreateSegmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`models.CreateSegment`](../../doc/models/create-segment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSegmentRequest := models.CreateSegmentRequest{
        Segment:              models.CreateSegment{
            SegmentProperty1Value: models.ToPointer(models.CreateSegmentSegmentProperty1ValueContainer.FromString("String1")),
            SegmentProperty2Value: models.ToPointer(models.CreateSegmentSegmentProperty2ValueContainer.FromString("String3")),
            SegmentProperty3Value: models.ToPointer(models.CreateSegmentSegmentProperty3ValueContainer.FromString("String1")),
            SegmentProperty4Value: models.ToPointer(models.CreateSegmentSegmentProperty4ValueContainer.FromString("String5")),
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
                models.CreateOrUpdateSegmentPrice{
                    StartingQuantity:     models.ToPointer(64),
                    EndingQuantity:       models.ToPointer(38),
                    UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
                },
            },
        },
    }

}
```

