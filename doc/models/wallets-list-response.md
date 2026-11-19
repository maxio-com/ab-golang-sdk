
# Wallets List Response

## Structure

`WalletsListResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wallets` | [`[]models.Wallet`](../../doc/models/wallet.md) | Required | - |
| `Meta` | [`models.ListPublicKeysMeta`](../../doc/models/list-public-keys-meta.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletsListResponse := models.WalletsListResponse{
        Wallets:              []models.Wallet{
            models.Wallet{
                Id:                         models.ToPointer(80),
                OwnerType:                  models.ToPointer(models.OwnerType_SUBSCRIPTION),
                OwnerId:                    models.ToPointer(156),
                Name:                       models.NewOptional(models.ToPointer("name4")),
                WalletProductId:            models.ToPointer(72),
                Unit:                       models.ToPointer("tokens"),
                Balance:                    models.ToPointer("150.0"),
                PostedBalance:              models.ToPointer("150.0"),
                DepletionPolicy:            models.ToPointer("hard_stop"),
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

