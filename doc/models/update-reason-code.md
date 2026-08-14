
# Update Reason Code

## Structure

`UpdateReasonCode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | The unique identifier for the ReasonCode |
| `Description` | `*string` | Optional | The friendly summary of what the code signifies |
| `Position` | `*int` | Optional | The order that code appears in lists |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateReasonCode := models.UpdateReasonCode{
        Code:                 models.ToPointer("code4"),
        Description:          models.ToPointer("description6"),
        Position:             models.ToPointer(4),
    }

}
```

