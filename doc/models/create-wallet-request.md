
# Create Wallet Request

## Structure

`CreateWalletRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wallet` | [`models.Wallet1`](../../doc/models/wallet-1.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createWalletRequest := models.CreateWalletRequest{
        Wallet:               models.Wallet1{
            SubscriptionId:          94,
            WalletProductId:         models.Wallet1WalletProductIdContainer.FromNumber(154),
            Name:                    models.NewOptional(models.ToPointer("name6")),
            DepletionPolicy:         models.ToPointer("depletion_policy4"),
            HardStopOverflow:        models.ToPointer(models.WalletHardStopOverflow_DISCARD),
            MeterId:                 models.ToPointer(146),
            ConversionRate:          models.ToPointer(models.Wallet1ConversionRateContainer.FromString("String9")),
            InitialBalance:          models.ToPointer(models.Wallet1InitialBalanceContainer.FromString("String1")),
        },
    }

}
```

