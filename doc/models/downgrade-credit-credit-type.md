
# Downgrade Credit Credit Type

The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. Values are:

`full` -  A full price credit is added for the amount owed.

`prorated` - A prorated credit is added for the amount owed.

`none` - No charge is added.

## Enumeration

`DowngradeCreditCreditType`

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
    downgradeCreditCreditType := models.DowngradeCreditCreditType_FULL

}
```

