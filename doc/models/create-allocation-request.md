
# Create Allocation Request

## Structure

`CreateAllocationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Allocation` | [`models.CreateAllocation`](../../doc/models/create-allocation.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createAllocationRequest := models.CreateAllocationRequest{
        Allocation:           models.CreateAllocation{
            Quantity:                 float64(228.94),
            DecimalQuantity:          models.ToPointer("decimal_quantity6"),
            PreviousQuantity:         models.ToPointer(float64(254.04)),
            DecimalPreviousQuantity:  models.ToPointer("decimal_previous_quantity8"),
            ComponentId:              models.ToPointer(8),
            Memo:                     models.ToPointer("memo2"),
        },
    }

}
```

