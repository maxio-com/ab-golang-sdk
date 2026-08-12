
# Component Price

## Structure

`ComponentPrice`

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
| `SegmentId` | `models.Optional[int]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPrice := models.ComponentPrice{
        Id:                   models.ToPointer(18),
        ComponentId:          models.ToPointer(128),
        StartingQuantity:     models.ToPointer(64),
        EndingQuantity:       models.NewOptional(models.ToPointer(218)),
        UnitPrice:            models.ToPointer("unit_price4"),
    }

}
```

