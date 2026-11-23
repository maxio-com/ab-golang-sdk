
# Create Wallet Transaction Request

## Structure

`CreateWalletTransactionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletTransaction` | [`models.WalletTransaction1`](../../doc/models/wallet-transaction-1.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createWalletTransactionRequest := models.CreateWalletTransactionRequest{
        WalletTransaction:    models.WalletTransaction1{
            Kind:                 models.Kind1_GRANT,
            Amount:               models.WalletTransaction1AmountContainer.FromString("String5"),
            Memo:                 models.NewOptional(models.ToPointer("memo6")),
            ExpiresAt:            models.NewOptional(models.ToPointer("expires_at6")),
        },
    }

}
```

