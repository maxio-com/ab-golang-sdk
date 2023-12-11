
# Currency Price

## Structure

`CurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `Price` | `*float64` | Optional | - |
| `FormattedPrice` | `*string` | Optional | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `Role` | [`*models.CurrencyPriceRoleEnum`](currency-price-role-enum.md) | Optional | Role for the price. |

## Example (as JSON)

```json
{
  "id": 88,
  "currency": "currency6",
  "price": 41.36,
  "formatted_price": "formatted_price4",
  "product_price_point_id": 210
}
```

