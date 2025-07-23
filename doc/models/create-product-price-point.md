
# Create Product Price Point

## Structure

`CreateProductPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | The product price point name |
| `Handle` | `*string` | Optional | The product price point API handle |
| `PriceInCents` | `int64` | Required | The product price point price, in integer cents |
| `Interval` | `int` | Required | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product price point would renew every 30 days |
| `IntervalUnit` | [`models.IntervalUnit`](../../doc/models/interval-unit.md) | Required | A string representing the interval unit for this product price point, either month or day |
| `TrialPriceInCents` | `*int64` | Optional | The product price point trial price, in integer cents |
| `TrialInterval` | `*int` | Optional | The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product price point trial would last 30 days. |
| `TrialIntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the trial interval unit for this product price point, either month or day |
| `TrialType` | `*string` | Optional | - |
| `InitialChargeInCents` | `*int64` | Optional | The product price point initial charge, in integer cents |
| `InitialChargeAfterTrial` | `*bool` | Optional | - |
| `ExpirationInterval` | `*int` | Optional | The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | A string representing the expiration interval unit for this product price point, either month, day or never |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether or not to use the site's exchange rate or define your own pricing when your site has multiple currencies defined.<br><br>**Default**: `true` |

## Example (as JSON)

```json
{
  "name": "name6",
  "price_in_cents": 216,
  "interval": 200,
  "interval_unit": "day",
  "use_site_exchange_rate": true,
  "handle": "handle2",
  "trial_price_in_cents": 48,
  "trial_interval": 102,
  "trial_interval_unit": "day",
  "trial_type": "trial_type0"
}
```

