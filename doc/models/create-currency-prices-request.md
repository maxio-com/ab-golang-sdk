
# Create Currency Prices Request

## Structure

`CreateCurrencyPricesRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.CreateCurrencyPrice`](create-currency-price.md) | Required | - |

## Example (as JSON)

```json
{
  "currency_prices": [
    {
      "currency": "currency8",
      "price": 78,
      "price_id": 116
    }
  ]
}
```

