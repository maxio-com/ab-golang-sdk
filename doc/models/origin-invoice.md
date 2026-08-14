
# Origin Invoice

## Structure

`OriginInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | The UID of the invoice serving as an origin invoice. |
| `Number` | `*string` | Optional | The number of the invoice serving as an origin invoice. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    originInvoice := models.OriginInvoice{
        Uid:                  models.ToPointer("uid8"),
        Number:               models.ToPointer("number4"),
    }

}
```

