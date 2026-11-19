
# Price 1

## Structure

`Price1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `models.Optional[int]` | Optional | - |
| `UnitPrice` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    price1 := models.Price1{
        StartingQuantity:     models.ToPointer(138),
        EndingQuantity:       models.NewOptional(models.ToPointer(112)),
        UnitPrice:            models.ToPointer("unit_price8"),
    }

}
```

