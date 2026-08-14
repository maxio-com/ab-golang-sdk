
# Component Price Point Error Item

## Structure

`ComponentPricePointErrorItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `Message` | `*string` | Optional | - |
| `PricePoint` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentPricePointErrorItem := models.ComponentPricePointErrorItem{
        ComponentId:          models.ToPointer(174),
        Message:              models.ToPointer("message2"),
        PricePoint:           models.ToPointer(72),
    }

}
```

