
# Component

## Structure

`Component`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The unique ID assigned to the component by Chargify. This ID can be used to fetch the component from the API. |
| `Name` | `*string` | Optional | The name of the Component, suitable for display on statements. i.e. Text Messages. |
| `Handle` | `Optional[string]` | Optional | The component API handle |
| `PricingScheme` | [`Optional[models.PricingSchemeEnum]`](pricing-scheme-enum.md) | Optional | - |
| `UnitName` | `*string` | Optional | The name of the unit that the component’s usage is measured in. i.e. message |
| `UnitPrice` | `Optional[string]` | Optional | The amount the customer will be charged per unit. This field is only populated for ‘per_unit’ pricing schemes, otherwise it may be null. |
| `ProductFamilyId` | `*int` | Optional | The id of the Product Family to which the Component belongs |
| `ProductFamilyName` | `*string` | Optional | The name of the Product Family to which the Component belongs |
| `PricePerUnitInCents` | `Optional[int64]` | Optional | deprecated - use unit_price instead |
| `Kind` | [`*models.ComponentKindEnum`](component-kind-enum.md) | Optional | A handle for the component type |
| `Archived` | `*bool` | Optional | Boolean flag describing whether a component is archived or not. |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `Description` | `Optional[string]` | Optional | The description of the component. |
| `DefaultPricePointId` | `*int` | Optional | - |
| `Prices` | [`Optional[[]models.ComponentPrice]`](component-price.md) | Optional | An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty. |
| `PricePointCount` | `*int` | Optional | Count for the number of price points associated with the component |
| `PricePointsUrl` | `*string` | Optional | URL that points to the location to read the existing price points via GET request |
| `DefaultPricePointName` | `*string` | Optional | - |
| `TaxCode` | `Optional[string]` | Optional | A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters. |
| `Recurring` | `*bool` | Optional | - |
| `UpgradeCharge` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `CreatedAt` | `*string` | Optional | Timestamp indicating when this component was created |
| `UpdatedAt` | `*string` | Optional | Timestamp indicating when this component was updated |
| `ArchivedAt` | `Optional[string]` | Optional | Timestamp indicating when this component was archived |
| `HideDateRangeOnInvoice` | `*bool` | Optional | (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices. |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `ItemCategory` | [`Optional[models.ItemCategoryEnum]`](item-category-enum.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |
| `UseSiteExchangeRate` | `Optional[bool]` | Optional | - |
| `AccountingCode` | `Optional[string]` | Optional | E.g. Internal ID or SKU Number |
| `EventBasedBillingMetricId` | `*int` | Optional | (Only for Event Based Components) This is an ID of a metric attached to the component. This metric is used to bill upon collected events. |

## Example (as JSON)

```json
{
  "item_category": "Business Software",
  "id": 24,
  "name": "name2",
  "handle": "handle8",
  "pricing_scheme": "per_unit",
  "unit_name": "unit_name4"
}
```

