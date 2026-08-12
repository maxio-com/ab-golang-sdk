
# Count Response

## Structure

`CountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    countResponse := models.CountResponse{
        Count:                models.ToPointer(56),
    }

}
```

