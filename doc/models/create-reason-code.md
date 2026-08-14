
# Create Reason Code

## Structure

`CreateReasonCode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | The unique identifier for the ReasonCode |
| `Description` | `string` | Required | The friendly summary of what the code signifies |
| `Position` | `*int` | Optional | The order that code appears in lists |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createReasonCode := models.CreateReasonCode{
        Code:                 "code4",
        Description:          "description6",
        Position:             models.ToPointer(40),
    }

}
```

