
# Subscription Custom Price

(Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription

## Structure

`SubscriptionCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | (Optional) |
| `Handle` | `*string` | Optional | (Optional) |
| `PriceInCents` | `interface{}` | Required | Required if using `custom_price` attribute. |
| `Interval` | `interface{}` | Required | Required if using `custom_price` attribute. |
| `IntervalUnit` | [`models.IntervalUnitEnum`](interval-unit-enum.md) | Required | Required if using `custom_price` attribute. |
| `TrialPriceInCents` | `*interface{}` | Optional | (Optional) |
| `TrialInterval` | `*interface{}` | Optional | (Optional) |
| `TrialIntervalUnit` | [`*models.IntervalUnitEnum`](interval-unit-enum.md) | Optional | (Optional) |
| `InitialChargeInCents` | `*interface{}` | Optional | (Optional) |
| `InitialChargeAfterTrial` | `*bool` | Optional | (Optional) |
| `ExpirationInterval` | `*interface{}` | Optional | (Optional) |
| `ExpirationIntervalUnit` | [`*models.IntervalUnitEnum`](interval-unit-enum.md) | Optional | (Optional) |
| `TaxIncluded` | `*bool` | Optional | (Optional) |

## Example (as JSON)

```json
{
  "name": "name4",
  "handle": "handle0",
  "price_in_cents": {
    "key1": "val1",
    "key2": "val2"
  },
  "interval": {
    "key1": "val1",
    "key2": "val2"
  },
  "interval_unit": "day",
  "trial_price_in_cents": {
    "key1": "val1",
    "key2": "val2"
  },
  "trial_interval": {
    "key1": "val1",
    "key2": "val2"
  },
  "trial_interval_unit": "day"
}
```

