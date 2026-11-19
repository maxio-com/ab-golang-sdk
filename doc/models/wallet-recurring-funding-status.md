
# Wallet Recurring Funding Status

`active`: refilling normally. `erroring`: the schedule is still due and retrying, but the last attempt couldn't be written (see `reason`); correcting the underlying entitlement resumes funding on the next run with no action needed on the wallet itself. `stopped`: funding has ended for good (see `reason`); units already granted stay on the wallet and remain spendable, but no further automatic grants will be made.

## Enumeration

`WalletRecurringFundingStatus`

## Fields

| Name |
|  --- |
| `ACTIVE` |
| `ERRORING` |
| `STOPPED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletRecurringFundingStatus := models.WalletRecurringFundingStatus_ACTIVE

}
```

