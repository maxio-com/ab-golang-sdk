
# List MRR Response Result

## Structure

`ListMRRResponseResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Page` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `TotalEntries` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `CurrencySymbol` | `*string` | Optional | - |
| `Movements` | [`[]models.Movement`](../../doc/models/movement.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMRRResponseResult := models.ListMRRResponseResult{
        Page:                 models.ToPointer(40),
        PerPage:              models.ToPointer(208),
        TotalPages:           models.ToPointer(82),
        TotalEntries:         models.ToPointer(78),
        Currency:             models.ToPointer("currency6"),
    }

}
```

