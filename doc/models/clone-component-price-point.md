
# Clone Component Price Point

## Structure

`CloneComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cloneComponentPricePoint := models.CloneComponentPricePoint{
        Name:                 "name4",
        Handle:               models.ToPointer("handle0"),
    }

}
```

