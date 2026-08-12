
# Update Component

## Structure

`UpdateComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Handle` | `*string` | Optional | - |
| `Name` | `*string` | Optional | The name of the Component, suitable for display on statements. e.g., Text Messages. |
| `Description` | `models.Optional[string]` | Optional | The description of the component. |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `TaxCode` | `models.Optional[string]` | Optional | A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `ItemCategory` | [`models.Optional[models.ItemCategory]`](../../doc/models/item-category.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `UnspscCode` | `models.Optional[string]` | Optional | (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this component instead of the default derived from item_category. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateComponent := models.UpdateComponent{
        Handle:               models.ToPointer("handle4"),
        Name:                 models.ToPointer("name8"),
        Description:          models.NewOptional(models.ToPointer("description8")),
        AccountingCode:       models.NewOptional(models.ToPointer("accounting_code4")),
        Taxable:              models.ToPointer(false),
        ItemCategory:         models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
    }

}
```

