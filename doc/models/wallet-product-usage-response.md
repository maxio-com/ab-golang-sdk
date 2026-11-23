
# Wallet Product Usage Response

## Structure

`WalletProductUsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletTemplates` | [`[]models.WalletProductUsageTemplate`](../../doc/models/wallet-product-usage-template.md) | Required | Wallet templates billing through this product at its default price point. |
| `PricePointOverrides` | [`[]models.WalletProductUsageTemplate`](../../doc/models/wallet-product-usage-template.md) | Required | Wallet templates billing through this product at a price point other than the default. |
| `ActiveWallets` | [`models.ActiveWallets`](../../doc/models/active-wallets.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductUsageResponse := models.WalletProductUsageResponse{
        WalletTemplates:      []models.WalletProductUsageTemplate{
            models.WalletProductUsageTemplate{
                PlanName:             models.NewOptional(models.ToPointer("plan_name4")),
                WalletTemplateName:   models.NewOptional(models.ToPointer("wallet_template_name2")),
                UnitName:             models.ToPointer("unit_name4"),
                PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism2")),
                PricePointName:       models.NewOptional(models.ToPointer("price_point_name4")),
            },
        },
        PricePointOverrides:  []models.WalletProductUsageTemplate{
            models.WalletProductUsageTemplate{
                PlanName:             models.NewOptional(models.ToPointer("plan_name0")),
                WalletTemplateName:   models.NewOptional(models.ToPointer("wallet_template_name2")),
                UnitName:             models.ToPointer("unit_name0"),
                PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism8")),
                PricePointName:       models.NewOptional(models.ToPointer("price_point_name8")),
            },
        },
        ActiveWallets:        models.ActiveWallets{
            Total:                164,
            Sample:               []models.WalletProductActiveWallet{
                models.WalletProductActiveWallet{
                    Id:                   models.ToPointer(180),
                    SubscriptionId:       models.ToPointer(34),
                    CompanyName:          models.ToPointer("company_name4"),
                    PurchaseMechanism:    models.NewOptional(models.ToPointer("purchase_mechanism8")),
                    PricePointName:       models.NewOptional(models.ToPointer("price_point_name8")),
                },
            },
        },
    }

}
```

