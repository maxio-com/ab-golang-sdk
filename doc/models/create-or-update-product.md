
# Create or Update Product

## Structure

`CreateOrUpdateProduct`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | The product name |
| `Handle` | `*string` | Optional | The product API handle |
| `Description` | `string` | Required | The product description |
| `AccountingCode` | `*string` | Optional | E.g. Internal ID or SKU Number |
| `RequireCreditCard` | `*bool` | Optional | Deprecated value that can be ignored unless you have legacy hosted pages. For Public Signup Page users, please read this attribute from under the signup page. |
| `PriceInCents` | `int64` | Required | The product price, in integer cents |
| `Interval` | `int` | Required | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product would renew every 30 days |
| `IntervalUnit` | [`models.IntervalUnit`](../../doc/models/interval-unit.md) | Required | A string representing the interval unit for this product, either month or day |
| `TrialPriceInCents` | `*int64` | Optional | The product trial price, in integer cents |
| `TrialInterval` | `*int` | Optional | The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product trial would last 30 days. |
| `TrialIntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the trial interval unit for this product, either month or day |
| `TrialType` | `*string` | Optional | - |
| `ExpirationInterval` | `*int` | Optional | The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product would expire after 30 days. |
| `ExpirationIntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the expiration interval unit for this product, either month or day |
| `AutoCreateSignupPage` | `*bool` | Optional | - |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the product type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.<br>**Constraints**: *Maximum Length*: `10` |

## Example (as JSON)

```json
{
  "name": "name8",
  "handle": "handle4",
  "description": "description8",
  "accounting_code": "accounting_code4",
  "require_credit_card": false,
  "price_in_cents": 190,
  "interval": 174,
  "interval_unit": "day",
  "trial_price_in_cents": 22,
  "trial_interval": 76
}
```

