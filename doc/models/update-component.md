
# Update Component

## Structure

`UpdateComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Handle` | `*string` | Optional | - |
| `Name` | `*string` | Optional | The name of the Component, suitable for display on statements. i.e. Text Messages. |
| `Description` | `models.Optional[string]` | Optional | The description of the component. |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `TaxCode` | `models.Optional[string]` | Optional | A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `ItemCategory` | [`models.Optional[models.ItemCategory]`](../../doc/models/item-category.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |

## Example (as JSON)

```json
{
  "item_category": "Business Software",
  "handle": "handle6",
  "name": "name0",
  "description": "description0",
  "accounting_code": "accounting_code6",
  "taxable": false
}
```

