
# Applied Credit Note Data

## Structure

`AppliedCreditNoteData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | The UID of the credit note |
| `Number` | `*string` | Optional | The number of the credit note |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    appliedCreditNoteData := models.AppliedCreditNoteData{
        Uid:                  models.ToPointer("uid4"),
        Number:               models.ToPointer("number2"),
    }

}
```

