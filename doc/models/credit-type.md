
# Credit Type

The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.

## Enumeration

`CreditType`

## Fields

| Name |
|  --- |
| `FULL` |
| `PRORATED` |
| `NONE` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    creditType := models.CreditType_PRORATED

}
```

