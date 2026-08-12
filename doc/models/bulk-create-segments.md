
# Bulk Create Segments

## Structure

`BulkCreateSegments`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | [`[]models.CreateSegment`](../../doc/models/create-segment.md) | Optional | **Constraints**: *Maximum Items*: `2000` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bulkCreateSegments := models.BulkCreateSegments{
        Segments:             []models.CreateSegment{
            models.CreateSegment{
                SegmentProperty1Value: models.ToPointer(models.CreateSegmentSegmentProperty1ValueContainer.FromString("String3")),
                SegmentProperty2Value: models.ToPointer(models.CreateSegmentSegmentProperty2ValueContainer.FromString("String5")),
                SegmentProperty3Value: models.ToPointer(models.CreateSegmentSegmentProperty3ValueContainer.FromString("String3")),
                SegmentProperty4Value: models.ToPointer(models.CreateSegmentSegmentProperty4ValueContainer.FromString("String7")),
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
        },
    }

}
```

