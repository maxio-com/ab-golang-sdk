
# Price

## Structure

`Price`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | [`models.PriceStartingQuantity`](../../doc/models/containers/price-starting-quantity.md) | Required | This is a container for one-of cases. |
| `EndingQuantity` | [`models.Optional[models.PriceEndingQuantity]`](../../doc/models/containers/price-ending-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitPrice` | [`models.PriceUnitPrice`](../../doc/models/containers/price-unit-price.md) | Required | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    price := models.Price{
        StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(132),
        EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(6))),
        UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(70.44)),
    }

}
```

