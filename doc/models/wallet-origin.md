
# Wallet Origin

How the wallet came to exist. `manual` for a wallet created directly with [Create Wallet](../../doc/controllers/wallets.md#create-wallet). `feature` for a wallet auto-provisioned because a subscription resolved a usage-limit entitlement whose Feature has `wallet_behavior` configured.

## Enumeration

`WalletOrigin`

## Fields

| Name |
|  --- |
| `MANUAL` |
| `FEATURE` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletOrigin := models.WalletOrigin_MANUAL

}
```

