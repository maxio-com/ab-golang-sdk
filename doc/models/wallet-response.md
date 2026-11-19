
# Wallet Response

## Structure

`WalletResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wallet` | [`models.Wallet`](../../doc/models/wallet.md) | Required | A balance, denominated in the unit of a Wallet Product, owned by a subscription. Grants and recurring funding schedules are embedded so a single read shows the full balance breakdown. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletResponse := models.WalletResponse{
        Wallet:               models.Wallet{
            Id:                         models.ToPointer(240),
            OwnerType:                  models.ToPointer(models.OwnerType_SUBSCRIPTION),
            OwnerId:                    models.ToPointer(60),
            Name:                       models.NewOptional(models.ToPointer("name6")),
            WalletProductId:            models.ToPointer(232),
            Unit:                       models.ToPointer("tokens"),
            Balance:                    models.ToPointer("150.0"),
            PostedBalance:              models.ToPointer("150.0"),
            DepletionPolicy:            models.ToPointer("hard_stop"),
        },
    }

}
```

