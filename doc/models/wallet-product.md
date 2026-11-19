
# Wallet Product

## Structure

`WalletProduct`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `UnitName` | `string` | Required | The name of the unit the Wallet Product's usage is measured in, for example `tokens`. |
| `Handle` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Taxable` | `*bool` | Optional | - |
| `TaxCode` | `models.Optional[string]` | Optional | - |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `ItemCategory` | [`models.Optional[models.ItemCategory]`](../../doc/models/item-category.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |
| `PricePoint` | [`*models.WalletProductPricePointAttributes`](../../doc/models/wallet-product-price-point-attributes.md) | Optional | Creates one price point for the Wallet Product in the same call. Ignored if `price_points` is also sent. |
| `PricePoints` | [`[]models.WalletProductPricePointAttributes`](../../doc/models/wallet-product-price-point-attributes.md) | Optional | Creates several price points for the Wallet Product in the same call, winning over `price_point` if both are sent. The first price point created becomes the Wallet Product's default. The whole call is one transaction: if any price point in the array fails validation, nothing is created, including the Wallet Product itself. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProduct := models.WalletProduct{
        Name:                 "name4",
        UnitName:             "unit_name6",
        Handle:               models.ToPointer("handle0"),
        Description:          models.NewOptional(models.ToPointer("description4")),
        Taxable:              models.ToPointer(false),
        TaxCode:              models.NewOptional(models.ToPointer("tax_code2")),
        AccountingCode:       models.NewOptional(models.ToPointer("accounting_code0")),
        ItemCategory:         models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
    }

}
```

