
# Allocation Response

## Structure

`AllocationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Allocation` | [`*models.Allocation`](../../doc/models/allocation.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    allocationResponse := models.AllocationResponse{
        Allocation:           models.ToPointer(models.Allocation{
            AllocationId:             models.ToPointer(238),
            ComponentId:              models.ToPointer(8),
            ComponentHandle:          models.NewOptional(models.ToPointer("component_handle8")),
            SubscriptionId:           models.ToPointer(8),
            Quantity:                 models.ToPointer(models.AllocationQuantityContainer.FromNumber(32)),
        }),
    }

}
```

