
# Debit Note Status

Current status of the debit note.

## Enumeration

`DebitNoteStatus`

## Fields

| Name |
|  --- |
| `OPEN` |
| `APPLIED` |
| `BANISHED` |
| `PAID` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    debitNoteStatus := models.DebitNoteStatus_OPEN

}
```

