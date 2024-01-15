
# Create Product Price Point Request

## Structure

`CreateProductPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CreateProductPricePoint`](../../doc/models/create-product-price-point.md) | Required | - |

## Example (as JSON)

```json
{
  "price_point": {
    "name": "name0",
    "price_in_cents": 196,
    "interval": 44,
    "interval_unit": "day",
    "use_site_exchange_rate": true,
    "handle": "handle6",
    "trial_price_in_cents": 108,
    "trial_interval": 202,
    "trial_interval_unit": "day",
    "trial_type": "trial_type4"
  }
}
```

