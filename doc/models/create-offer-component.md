
# Create Offer Component

## Structure

`CreateOfferComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `StartingQuantity` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOfferComponent := models.CreateOfferComponent{
        ComponentId:          models.ToPointer(252),
        PricePointId:         models.ToPointer(20),
        StartingQuantity:     models.ToPointer(196),
    }

}
```

