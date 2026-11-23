
# Component Price Point Assignment

## Structure

`ComponentPricePointAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | [`*models.ComponentPricePointAssignmentComponentId`](../../doc/models/containers/component-price-point-assignment-component-id.md) | Optional | This is a container for one-of cases. |
| `PricePoint` | [`*models.ComponentPricePointAssignmentPricePoint`](../../doc/models/containers/component-price-point-assignment-price-point.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPricePointAssignment := models.ComponentPricePointAssignment{
        ComponentId:          models.ToPointer(models.ComponentPricePointAssignmentComponentIdContainer.FromNumber(112)),
        PricePoint:           models.ToPointer(models.ComponentPricePointAssignmentPricePointContainer.FromString("String7")),
    }

}
```

