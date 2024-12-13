
# Prepaid Usage Component

## Structure

`PrepaidUsageComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes". |
| `UnitName` | `string` | Required | The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. i.e. “message”, which may then be shown as “5 messages” on a subscription’s component line-item |
| `Description` | `*string` | Optional | A description for the component that will be displayed to the user on the hosted signup page. |
| `Handle` | `*string` | Optional | A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.<br>**Constraints**: *Pattern*: `^[a-z0-9][a-z0-9\-_:.]*$` |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261149166733-Component-Pricing-Schemes#price-bracket-rules) for an overview of how price brackets work for different pricing schemes. |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `PricePoints` | [`[]models.CreatePrepaidUsageComponentPricePoint`](../../doc/models/create-prepaid-usage-component-price-point.md) | Optional | - |
| `UnitPrice` | [`*models.PrepaidUsageComponentUnitPrice`](../../doc/models/containers/prepaid-usage-component-unit-price.md) | Optional | This is a container for one-of cases. |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters. |
| `HideDateRangeOnInvoice` | `*bool` | Optional | (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices. |
| `OveragePricing` | [`models.OveragePricing`](../../doc/models/overage-pricing.md) | Required | - |
| `RolloverPrepaidRemainder` | `*bool` | Optional | Boolean which controls whether or not remaining units should be rolled over to the next period |
| `RenewPrepaidAllocation` | `*bool` | Optional | Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period |
| `ExpirationInterval` | `*float64` | Optional | (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | - |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `PublicSignupPageIds` | `[]int` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "unit_name": "unit_name4",
  "description": "description8",
  "handle": "handle8",
  "taxable": false,
  "pricing_scheme": "stairstep",
  "prices": [
    {
      "starting_quantity": 242,
      "ending_quantity": 40,
      "unit_price": 23.26
    },
    {
      "starting_quantity": 242,
      "ending_quantity": 40,
      "unit_price": 23.26
    }
  ],
  "upgrade_charge": "prorated",
  "overage_pricing": {
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ]
  }
}
```

