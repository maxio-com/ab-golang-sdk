
# Create Component Price Point Request

## Structure

`CreateComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CreateComponentPricePointRequestPricePoint`](../../doc/models/containers/create-component-price-point-request-price-point.md) | Required | This is a container for any-of cases. |

## Example (as JSON)

```json
{
  "price_point": {
    "name": "name0",
    "pricing_scheme": "per_unit",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ],
    "use_site_exchange_rate": true,
    "handle": "handle6",
    "tax_included": false,
    "interval": 24,
    "interval_unit": "day"
  }
}
```

