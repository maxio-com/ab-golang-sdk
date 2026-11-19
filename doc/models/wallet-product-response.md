
# Wallet Product Response

A Wallet Product is a Component created with no product family, so it returns the full Component field set under a `wallet_product` root instead of `component`.

## Structure

`WalletProductResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletProduct` | [`models.Component`](../../doc/models/component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductResponse := models.WalletProductResponse{
        WalletProduct:        models.Component{
            Id:                        models.ToPointer(224),
            Name:                      models.ToPointer("name0"),
            Handle:                    models.NewOptional(models.ToPointer("handle6")),
            PricingScheme:             models.NewOptional(models.ToPointer(models.PricingScheme_PERUNIT)),
            UnitName:                  models.ToPointer("unit_name2"),
            ItemCategory:              models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
        },
    }

}
```

