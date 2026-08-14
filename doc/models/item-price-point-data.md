
# Item Price Point Data

## Structure

`ItemPricePointData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    itemPricePointData := models.ItemPricePointData{
        Id:                   models.ToPointer(80),
        Handle:               models.ToPointer("handle8"),
        Name:                 models.ToPointer("name2"),
    }

}
```

