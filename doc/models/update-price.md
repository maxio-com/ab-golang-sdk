
# Update Price

## Structure

`UpdatePrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `EndingQuantity` | [`*models.UpdatePriceEndingQuantity`](../../doc/models/containers/update-price-ending-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitPrice` | [`*models.UpdatePriceUnitPrice`](../../doc/models/containers/update-price-unit-price.md) | Optional | This is a container for one-of cases. |
| `Destroy` | `*bool` | Optional | - |
| `StartingQuantity` | [`*models.UpdatePriceStartingQuantity`](../../doc/models/containers/update-price-starting-quantity.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updatePrice := models.UpdatePrice{
        Id:                   models.ToPointer(206),
        EndingQuantity:       models.ToPointer(models.UpdatePriceEndingQuantityContainer.FromNumber(28)),
        UnitPrice:            models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(181.3))),
        Destroy:              models.ToPointer(false),
        StartingQuantity:     models.ToPointer(models.UpdatePriceStartingQuantityContainer.FromNumber(54)),
    }

}
```

