
# Wallet Grant Result Response

Returned by an adjustment: the ledger row written, plus the wallet's new balance.

## Structure

`WalletGrantResultResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletTransactions` | [`[]models.WalletTransaction`](../../doc/models/wallet-transaction.md) | Required | - |
| `Balance` | `string` | Required | The wallet's new balance, as a decimal string. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletGrantResultResponse := models.WalletGrantResultResponse{
        WalletTransactions:   []models.WalletTransaction{
            models.WalletTransaction{
                Id:                     models.ToPointer(140),
                WalletId:               models.ToPointer(244),
                Kind:                   models.ToPointer("grant"),
                Amount:                 models.ToPointer("-30.0"),
                GrantId:                models.NewOptional(models.ToPointer(160)),
            },
        },
        Balance:              "balance8",
    }

}
```

