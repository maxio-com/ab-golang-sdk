
# Wallet Transaction 1

## Structure

`WalletTransaction1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Kind` | [`models.Kind1`](../../doc/models/kind-1.md) | Required | Determines the direction of the movement. The amount itself is always sent positive. |
| `Amount` | [`models.WalletTransaction1Amount`](../../doc/models/containers/wallet-transaction-1-amount.md) | Required | This is a container for one-of cases. |
| `Memo` | `models.Optional[string]` | Optional | **Constraints**: *Maximum Length*: `255` |
| `ExpiresAt` | `models.Optional[string]` | Optional | An expiration for the grant this creates. Only accepted with `kind: grant`; sending it with `kind: debit` is rejected. Omit it for an evergreen grant. Accepts either a bare calendar date (`YYYY-MM-DD`), resolved to the end of that day in the site's time zone, or a full timestamp. A calendar day that doesn't exist, such as `2027-02-30`, is rejected with `422` rather than rolled into the next month. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletTransaction1 := models.WalletTransaction1{
        Kind:                 models.Kind1_GRANT,
        Amount:               models.WalletTransaction1AmountContainer.FromString("String5"),
        Memo:                 models.NewOptional(models.ToPointer("memo6")),
        ExpiresAt:            models.NewOptional(models.ToPointer("expires_at4")),
    }

}
```

