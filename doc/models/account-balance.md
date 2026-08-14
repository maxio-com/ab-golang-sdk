
# Account Balance

## Structure

`AccountBalance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BalanceInCents` | `*int64` | Optional | The balance in cents. |
| `AutomaticBalanceInCents` | `models.Optional[int64]` | Optional | The automatic balance in cents. |
| `RemittanceBalanceInCents` | `models.Optional[int64]` | Optional | The remittance balance in cents. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    accountBalance := models.AccountBalance{
        BalanceInCents:           models.ToPointer(int64(242)),
        AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(0))),
        RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(32))),
    }

}
```

