
# Segment Response

## Structure

`SegmentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`*models.Segment`](../../doc/models/segment.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    segmentResponse := models.SegmentResponse{
        Segment:              models.ToPointer(models.Segment{
            Id:                        models.ToPointer(118),
            ComponentId:               models.ToPointer(228),
            PricePointId:              models.ToPointer(4),
            EventBasedBillingMetricId: models.ToPointer(56),
            PricingScheme:             models.ToPointer(models.PricingScheme_STAIRSTEP),
        }),
    }

}
```

