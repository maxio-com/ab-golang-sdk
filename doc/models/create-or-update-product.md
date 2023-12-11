
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
| `IntervalUnit` | [`models.IntervalUnitEnum`](interval-unit-enum.md) | Required | A string representing the interval unit for this product, either month or day |
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
  "auto_create_signup_page": false,
  "tax_code": "tax_code6"
}
```

