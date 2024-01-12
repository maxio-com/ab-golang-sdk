
# Price Point

## Structure

`PricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | - |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.<br>**Default**: `true` |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `OveragePricing` | [`*models.OveragePricing`](../../doc/models/overage-pricing.md) | Optional | - |
| `RolloverPrepaidRemainder` | `*bool` | Optional | Boolean which controls whether or not remaining units should be rolled over to the next period |
| `RenewPrepaidAllocation` | `*bool` | Optional | Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period |
| `ExpirationInterval` | `*float64` | Optional | (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire |
| `ExpirationIntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | - |

## Example (as JSON)

```json
{
  "use_site_exchange_rate": true,
  "name": "name0",
  "handle": "handle6",
  "pricing_scheme": "per_unit",
  "prices": [
    {
      "starting_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "ending_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ]
}
```

