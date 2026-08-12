
# List Segments Response

## Structure

`ListSegmentsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | [`[]models.Segment`](../../doc/models/segment.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSegmentsResponse := models.ListSegmentsResponse{
        Segments:             []models.Segment{
            models.Segment{
                Id:                        models.ToPointer(50),
                ComponentId:               models.ToPointer(160),
                PricePointId:              models.ToPointer(184),
                EventBasedBillingMetricId: models.ToPointer(244),
                PricingScheme:             models.ToPointer(models.PricingScheme_STAIRSTEP),
            },
            models.Segment{
                Id:                        models.ToPointer(50),
                ComponentId:               models.ToPointer(160),
                PricePointId:              models.ToPointer(184),
                EventBasedBillingMetricId: models.ToPointer(244),
                PricingScheme:             models.ToPointer(models.PricingScheme_STAIRSTEP),
            },
        },
    }

}
```

