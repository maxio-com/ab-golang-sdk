
# Segment Price

## Structure

`SegmentPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `models.Optional[int]` | Optional | - |
| `UnitPrice` | `*string` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `FormattedUnitPrice` | `*string` | Optional | - |
| `SegmentId` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    segmentPrice := models.SegmentPrice{
        Id:                   models.ToPointer(194),
        ComponentId:          models.ToPointer(48),
        StartingQuantity:     models.ToPointer(144),
        EndingQuantity:       models.NewOptional(models.ToPointer(118)),
        UnitPrice:            models.ToPointer("unit_price0"),
    }

}
```

