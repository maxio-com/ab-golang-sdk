
# List Public Keys Meta

## Structure

`ListPublicKeysMeta`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalCount` | `*int` | Optional | - |
| `CurrentPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listPublicKeysMeta := models.ListPublicKeysMeta{
        TotalCount:           models.ToPointer(232),
        CurrentPage:          models.ToPointer(208),
        TotalPages:           models.ToPointer(220),
        PerPage:              models.ToPointer(70),
    }

}
```

