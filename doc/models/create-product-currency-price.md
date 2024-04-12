
# Create Product Currency Price

## Structure

`CreateProductCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currency` | `string` | Required | ISO code for one of the site level currencies. |
| `Price` | `int` | Required | Price for the given role. |
| `Role` | [`models.CurrencyPriceRole`](../../doc/models/currency-price-role.md) | Required | Role for the price. |

## Example (as JSON)

```json
{
  "currency": "currency0",
  "price": 222,
  "role": "baseline"
}
```

