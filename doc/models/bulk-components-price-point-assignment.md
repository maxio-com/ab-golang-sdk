
# Bulk Components Price Point Assignment

## Structure

`BulkComponentsPricePointAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Components` | [`[]models.ComponentPricePointAssignment`](../../doc/models/component-price-point-assignment.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bulkComponentsPricePointAssignment := models.BulkComponentsPricePointAssignment{
        Components:           []models.ComponentPricePointAssignment{
            models.ComponentPricePointAssignment{
                ComponentId:          models.ToPointer(108),
                PricePoint:           models.ToPointer(models.ComponentPricePointAssignmentPricePointContainer.FromString("String5")),
            },
        },
    }

}
```

