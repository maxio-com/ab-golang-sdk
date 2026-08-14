
# Create Reason Code Request

## Structure

`CreateReasonCodeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ReasonCode` | [`models.CreateReasonCode`](../../doc/models/create-reason-code.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createReasonCodeRequest := models.CreateReasonCodeRequest{
        ReasonCode:           models.CreateReasonCode{
            Code:                 "code4",
            Description:          "description6",
            Position:             models.ToPointer(14),
        },
    }

}
```

