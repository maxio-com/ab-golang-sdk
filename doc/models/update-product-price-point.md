
# Update Product Price Point

## Structure

`UpdateProductPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Handle` | `*string` | Optional | - |
| `PriceInCents` | `*int64` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateProductPricePoint := models.UpdateProductPricePoint{
        Handle:               models.ToPointer("handle2"),
        PriceInCents:         models.ToPointer(int64(154)),
    }

}
```

