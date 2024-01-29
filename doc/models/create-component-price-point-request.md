
# Create Component Price Point Request

## Structure

`CreateComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.PricePoint`](../../doc/models/price-point.md) | Required | - |

## Example (as JSON)

```json
{
  "price_point": {
    "use_site_exchange_rate": true,
    "name": "name0",
    "handle": "handle6",
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
      },
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
      },
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
}
```

