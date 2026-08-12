
# Metafield Input

Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected. Defaults to 'text'.

## Enumeration

`MetafieldInput`

## Fields

| Name |
|  --- |
| `BALANCETRACKER` |
| `TEXT` |
| `RADIO` |
| `DROPDOWN` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    metafieldInput := models.MetafieldInput_BALANCETRACKER

}
```

