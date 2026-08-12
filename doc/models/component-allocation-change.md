
# Component Allocation Change

## Structure

`ComponentAllocationChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousAllocation` | `int` | Required | - |
| `NewAllocation` | `int` | Required | - |
| `ComponentId` | `int` | Required | - |
| `ComponentHandle` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `AllocationId` | `int` | Required | - |
| `AllocatedQuantity` | [`*models.ComponentAllocationChangeAllocatedQuantity`](../../doc/models/containers/component-allocation-change-allocated-quantity.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentAllocationChange := models.ComponentAllocationChange{
        PreviousAllocation:   78,
        NewAllocation:        118,
        ComponentId:          72,
        ComponentHandle:      "component_handle8",
        Memo:                 "memo2",
        AllocationId:         174,
        AllocatedQuantity:    models.ToPointer(models.ComponentAllocationChangeAllocatedQuantityContainer.FromNumber(88)),
    }

}
```

