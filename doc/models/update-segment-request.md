
# Update Segment Request

## Structure

`UpdateSegmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`models.UpdateSegment`](../../doc/models/update-segment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSegmentRequest := models.UpdateSegmentRequest{
        Segment:              models.UpdateSegment{
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
    }

}
```

