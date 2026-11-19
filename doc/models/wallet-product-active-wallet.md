
# Wallet Product Active Wallet

One active wallet drawing on this Wallet Product.

## Structure

`WalletProductActiveWallet`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Advanced Billing id of the wallet. |
| `SubscriptionId` | `*int` | Optional | - |
| `CompanyName` | `*string` | Optional | The owning customer's organization name, falling back to their full name, then to `"Customer #<id>"`. `"Customer"` (with no id) when the wallet's subscription has no customer at all. |
| `PurchaseMechanism` | `models.Optional[string]` | Optional | See the same field on [Wallet Product Usage Template](../../doc/models/wallet-product-usage-template.md). The same caveats apply. |
| `PricePointName` | `models.Optional[string]` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `Active` | `*bool` | Optional | `false` when the wallet's Feature is archived. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductActiveWallet := models.WalletProductActiveWallet{
        Id:                   models.ToPointer(214),
        SubscriptionId:       models.ToPointer(68),
        CompanyName:          models.ToPointer("company_name8"),
        PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism2")),
        PricePointName:       models.NewOptional(models.ToPointer("price_point_name4")),
    }

}
```

