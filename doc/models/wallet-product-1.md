
# Wallet Product 1

## Structure

`WalletProduct1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Taxable` | `*bool` | Optional | - |
| `TaxCode` | `models.Optional[string]` | Optional | - |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `ItemCategory` | [`models.Optional[models.ItemCategory]`](../../doc/models/item-category.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProduct1 := models.WalletProduct1{
        Name:                 models.ToPointer("name8"),
        Description:          models.NewOptional(models.ToPointer("description8")),
        Taxable:              models.ToPointer(false),
        TaxCode:              models.NewOptional(models.ToPointer("tax_code6")),
        AccountingCode:       models.NewOptional(models.ToPointer("accounting_code4")),
        ItemCategory:         models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
    }

}
```

