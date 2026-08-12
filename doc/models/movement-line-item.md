
# Movement Line Item

## Structure

`MovementLineItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductId` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | For Product (or "baseline") line items, this field will have a value of `0`. |
| `PricePointId` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Mrr` | `*int` | Optional | - |
| `MrrMovements` | [`[]models.MRRMovement`](../../doc/models/mrr-movement.md) | Optional | - |
| `Quantity` | `*int` | Optional | - |
| `PrevQuantity` | `*int` | Optional | - |
| `Recurring` | `*bool` | Optional | When `true`, the line item's MRR value will contribute to the `plan` breakout. When `false`, the line item contributes to the `usage` breakout. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    movementLineItem := models.MovementLineItem{
        ProductId:            models.ToPointer(146),
        ComponentId:          models.ToPointer(58),
        PricePointId:         models.ToPointer(82),
        Name:                 models.ToPointer("name8"),
        Mrr:                  models.ToPointer(92),
    }

}
```

