
# Update Reason Code Request

## Structure

`UpdateReasonCodeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ReasonCode` | [`models.UpdateReasonCode`](../../doc/models/update-reason-code.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateReasonCodeRequest := models.UpdateReasonCodeRequest{
        ReasonCode:           models.UpdateReasonCode{
            Code:                 models.ToPointer("code4"),
            Description:          models.ToPointer("description6"),
            Position:             models.ToPointer(14),
        },
    }

}
```

