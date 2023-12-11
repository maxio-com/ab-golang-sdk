
# Product Price Point Error Response Exception

## Structure

`ProductPricePointErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.ProductPricePointErrors`](product-price-point-errors.md) | Required | - |

## Example (as JSON)

```json
{
  "errors": {
    "price_point": "can't be blank",
    "interval": [
      "Recurring Interval: cannot be blank.",
      "Recurring Interval: must be greater than or equal to 1."
    ],
    "interval_unit": [
      "Interval unit: cannot be blank.",
      "Interval unit: must be 'month' or 'day'."
    ],
    "name": [
      "Name: cannot be blank."
    ],
    "price": [
      "Price: is not a number.",
      "Price: must be greater than or equal to 0."
    ],
    "price_in_cents": [
      "Price in cents: cannot be blank."
    ]
  }
}
```

