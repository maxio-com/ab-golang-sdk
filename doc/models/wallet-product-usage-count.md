
# Wallet Product Usage Count

A rollup of what's using this Wallet Product, for the catalog listing.

## Structure

`WalletProductUsageCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletConfigurations` | `int` | Required | Number of wallet templates billing through this product, whether or not any wallet has been created from them yet. |
| `ActiveWallets` | `int` | Required | Number of active (non-archived) wallets currently denominated by this product. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductUsageCount := models.WalletProductUsageCount{
        WalletConfigurations: 154,
        ActiveWallets:        72,
    }

}
```

