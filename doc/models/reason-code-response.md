
# Reason Code Response

## Structure

`ReasonCodeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ReasonCode` | [`models.ReasonCode`](../../doc/models/reason-code.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reasonCodeResponse := models.ReasonCodeResponse{
        ReasonCode:           models.ReasonCode{
            Id:                   models.ToPointer(240),
            SiteId:               models.ToPointer(166),
            Code:                 models.ToPointer("code4"),
            Description:          models.ToPointer("description6"),
            Position:             models.ToPointer(14),
        },
    }

}
```

