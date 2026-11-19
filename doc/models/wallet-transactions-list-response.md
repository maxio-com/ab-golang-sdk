
# Wallet Transactions List Response

## Structure

`WalletTransactionsListResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletTransactions` | [`[]models.WalletTransaction`](../../doc/models/wallet-transaction.md) | Required | - |
| `Meta` | [`models.ListPublicKeysMeta`](../../doc/models/list-public-keys-meta.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletTransactionsListResponse := models.WalletTransactionsListResponse{
        WalletTransactions:   []models.WalletTransaction{
            models.WalletTransaction{
                Id:                     models.ToPointer(140),
                WalletId:               models.ToPointer(244),
                Kind:                   models.ToPointer("grant"),
                Amount:                 models.ToPointer("-30.0"),
                GrantId:                models.NewOptional(models.ToPointer(160)),
            },
        },
        Meta:                 models.ListPublicKeysMeta{
            TotalCount:           models.ToPointer(150),
            CurrentPage:          models.ToPointer(126),
            TotalPages:           models.ToPointer(138),
            PerPage:              models.ToPointer(152),
        },
    }

}
```

