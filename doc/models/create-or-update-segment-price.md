
# Create or Update Segment Price

## Structure

`CreateOrUpdateSegmentPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `*int` | Optional | - |
| `UnitPrice` | [`models.CreateOrUpdateSegmentPriceUnitPrice`](../../doc/models/containers/create-or-update-segment-price-unit-price.md) | Required | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOrUpdateSegmentPrice := models.CreateOrUpdateSegmentPrice{
        StartingQuantity:     models.ToPointer(98),
        EndingQuantity:       models.ToPointer(184),
        UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String9"),
    }

}
```

