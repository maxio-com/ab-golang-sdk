
# Active Wallets

## Structure

`ActiveWallets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Total` | `int` | Required | Total count of active wallets drawing on this product. |
| `Sample` | [`[]models.WalletProductActiveWallet`](../../doc/models/wallet-product-active-wallet.md) | Required | Up to 50 of the most recently created active wallets, newest first. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    activeWallets := models.ActiveWallets{
        Total:                120,
        Sample:               []models.WalletProductActiveWallet{
            models.WalletProductActiveWallet{
                Id:                   models.ToPointer(180),
                SubscriptionId:       models.ToPointer(34),
                CompanyName:          models.ToPointer("company_name4"),
                PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism8")),
                PricePointName:       models.NewOptional(models.ToPointer("price_point_name8")),
            },
        },
    }

}
```

