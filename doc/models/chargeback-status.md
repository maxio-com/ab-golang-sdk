
# Chargeback Status

The current chargeback status.

## Enumeration

`ChargebackStatus`

## Fields

| Name |
|  --- |
| `OPEN` |
| `LOST` |
| `WON` |
| `CLOSED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    chargebackStatus := models.ChargebackStatus_OPEN

}
```

