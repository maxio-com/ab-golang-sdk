
# Component Kind

A handle for the component type

## Enumeration

`ComponentKind`

## Fields

| Name |
|  --- |
| `METEREDCOMPONENT` |
| `QUANTITYBASEDCOMPONENT` |
| `ONOFFCOMPONENT` |
| `PREPAIDUSAGECOMPONENT` |
| `EVENTBASEDCOMPONENT` |
| `WALLETPRODUCT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentKind := models.ComponentKind_EVENTBASEDCOMPONENT

}
```

