
# Update Component Price Point Request

## Structure

`UpdateComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`*models.UpdateComponentPricePoint`](../../doc/models/update-component-price-point.md) | Optional | - |

## Example (as JSON)

```json
{
  "price_point": {
    "name": "name0",
    "interval": 44,
    "interval_unit": "day",
    "prices": [
      {
        "id": 18,
        "ending_quantity": 38,
        "unit_price": 88,
        "_destroy": "_destroy4",
        "starting_quantity": 64
      },
      {
        "id": 18,
        "ending_quantity": 38,
        "unit_price": 88,
        "_destroy": "_destroy4",
        "starting_quantity": 64
      },
      {
        "id": 18,
        "ending_quantity": 38,
        "unit_price": 88,
        "_destroy": "_destroy4",
        "starting_quantity": 64
      }
    ]
  }
}
```

