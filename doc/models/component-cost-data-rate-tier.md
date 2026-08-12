
# Component Cost Data Rate Tier

## Structure

`ComponentCostDataRateTier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `models.Optional[int]` | Optional | - |
| `Quantity` | `*string` | Optional | - |
| `UnitPrice` | `*string` | Optional | - |
| `Amount` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentCostDataRateTier := models.ComponentCostDataRateTier{
        StartingQuantity:     models.ToPointer(204),
        EndingQuantity:       models.NewOptional(models.ToPointer(178)),
        Quantity:             models.ToPointer("quantity4"),
        UnitPrice:            models.ToPointer("unit_price6"),
        Amount:               models.ToPointer("amount0"),
    }

}
```

