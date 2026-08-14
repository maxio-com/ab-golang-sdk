
# Change Chargeback Status Event Data

Example schema for an `change_chargeback_status` event

## Structure

`ChangeChargebackStatusEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargebackStatus` | [`models.ChargebackStatus`](../../doc/models/chargeback-status.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    changeChargebackStatusEventData := models.ChangeChargebackStatusEventData{
        ChargebackStatus:     models.ChargebackStatus_WON,
    }

}
```

