
# Bulk Update Segments

## Structure

`BulkUpdateSegments`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | [`[]models.BulkUpdateSegmentsItem`](../../doc/models/bulk-update-segments-item.md) | Optional | **Constraints**: *Maximum Items*: `1000` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bulkUpdateSegments := models.BulkUpdateSegments{
        Segments:             []models.BulkUpdateSegmentsItem{
            models.BulkUpdateSegmentsItem{
                Id:                   50,
                PricingScheme:        models.PricingScheme_STAIRSTEP,
                Prices:               []models.CreateOrUpdateSegmentPrice{
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
            models.BulkUpdateSegmentsItem{
                Id:                   50,
                PricingScheme:        models.PricingScheme_STAIRSTEP,
                Prices:               []models.CreateOrUpdateSegmentPrice{
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

