
# Create Component Price Points Request

## Structure

`CreateComponentPricePointsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.PricePoint`](../../doc/models/price-point.md) | Required | - |

## Example (as JSON)

```json
{
  "price_points": [
    {
      "use_site_exchange_rate": true,
      "name": "name2",
      "handle": "handle8",
      "pricing_scheme": "per_unit",
      "prices": [
        {
          "starting_quantity": {
            "key1": "val1",
            "key2": "val2"
          },
          "ending_quantity": {
            "key1": "val1",
            "key2": "val2"
          },
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ]
    }
  ]
}
```

