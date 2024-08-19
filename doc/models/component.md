
# Component

## Structure

`Component`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The unique ID assigned to the component by Chargify. This ID can be used to fetch the component from the API. |
| `Name` | `*string` | Optional | The name of the Component, suitable for display on statements. i.e. Text Messages. |
| `Handle` | `models.Optional[string]` | Optional | The component API handle |
| `PricingScheme` | [`models.Optional[models.PricingScheme]`](../../doc/models/pricing-scheme.md) | Optional | - |
| `UnitName` | `*string` | Optional | The name of the unit that the component’s usage is measured in. i.e. message |
| `UnitPrice` | `models.Optional[string]` | Optional | The amount the customer will be charged per unit. This field is only populated for ‘per_unit’ pricing schemes, otherwise it may be null. |
| `ProductFamilyId` | `*int` | Optional | The id of the Product Family to which the Component belongs |
| `ProductFamilyName` | `*string` | Optional | The name of the Product Family to which the Component belongs |
| `PricePerUnitInCents` | `models.Optional[int64]` | Optional | deprecated - use unit_price instead |
| `Kind` | [`*models.ComponentKind`](../../doc/models/component-kind.md) | Optional | A handle for the component type |
| `Archived` | `*bool` | Optional | Boolean flag describing whether a component is archived or not. |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `Description` | `models.Optional[string]` | Optional | The description of the component. |
| `DefaultPricePointId` | `models.Optional[int]` | Optional | - |
| `OveragePrices` | [`models.Optional[[]models.ComponentPrice]`](../../doc/models/component-price.md) | Optional | An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty. |
| `Prices` | [`models.Optional[[]models.ComponentPrice]`](../../doc/models/component-price.md) | Optional | An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty. |
| `PricePointCount` | `*int` | Optional | Count for the number of price points associated with the component |
| `PricePointsUrl` | `models.Optional[string]` | Optional | URL that points to the location to read the existing price points via GET request |
| `DefaultPricePointName` | `*string` | Optional | - |
| `TaxCode` | `models.Optional[string]` | Optional | A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters. |
| `Recurring` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `CreatedAt` | `*time.Time` | Optional | Timestamp indicating when this component was created |
| `UpdatedAt` | `*time.Time` | Optional | Timestamp indicating when this component was updated |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | Timestamp indicating when this component was archived |
| `HideDateRangeOnInvoice` | `*bool` | Optional | (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices. |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `ItemCategory` | [`models.Optional[models.ItemCategory]`](../../doc/models/item-category.md) | Optional | One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other |
| `UseSiteExchangeRate` | `models.Optional[bool]` | Optional | - |
| `AccountingCode` | `models.Optional[string]` | Optional | E.g. Internal ID or SKU Number |
| `EventBasedBillingMetricId` | `*int` | Optional | (Only for Event Based Components) This is an ID of a metric attached to the component. This metric is used to bill upon collected events. |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled. |

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

