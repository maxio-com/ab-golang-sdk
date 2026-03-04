
# Scheduled Renewal Product Price Point

Custom pricing for a product within a scheduled renewal.

## Structure

`ScheduledRenewalProductPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | (Optional) |
| `Handle` | `*string` | Optional | (Optional) |
| `PriceInCents` | [`models.ScheduledRenewalProductPricePointPriceInCents`](../../doc/models/containers/scheduled-renewal-product-price-point-price-in-cents.md) | Required | This is a container for one-of cases. |
| `Interval` | [`models.ScheduledRenewalProductPricePointInterval`](../../doc/models/containers/scheduled-renewal-product-price-point-interval.md) | Required | This is a container for one-of cases. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Required | Required if using `custom_price` attribute. |
| `TaxIncluded` | `*bool` | Optional | (Optional) |
| `InitialChargeInCents` | `*int64` | Optional | The product price point initial charge, in integer cents. |
| `ExpirationInterval` | `*int` | Optional | The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | A string representing the expiration interval unit for this product price point, either month, day or never |

## Example (as JSON)

```json
{
  "name": "name4",
  "handle": "handle0",
  "price_in_cents": "String3",
  "interval": "String9",
  "interval_unit": "day",
  "tax_included": false,
  "initial_charge_in_cents": 86,
  "expiration_interval": 108
}
```

