
# Create Wallet Product Request

Does not accept `product_family_id`: a Wallet Product never belongs to a product family. The underlying endpoint accepts the param, but only to reject it; sending one returns `422` ("Product family cannot be set on a wallet product") rather than silently ignoring it.

## Structure

`CreateWalletProductRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletProduct` | [`models.WalletProduct`](../../doc/models/wallet-product.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createWalletProductRequest := models.CreateWalletProductRequest{
        WalletProduct:        models.WalletProduct{
            Name:                 "name0",
            UnitName:             "unit_name2",
            Handle:               models.ToPointer("handle6"),
            Description:          models.NewOptional(models.ToPointer("description0")),
            Taxable:              models.ToPointer(false),
            TaxCode:              models.NewOptional(models.ToPointer("tax_code8")),
            AccountingCode:       models.NewOptional(models.ToPointer("accounting_code6")),
            ItemCategory:         models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
        },
    }

}
```

