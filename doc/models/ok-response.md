
# Ok Response

## Structure

`OkResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ok` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    okResponse := models.OkResponse{
        Ok:                   models.ToPointer("ok8"),
    }

}
```

