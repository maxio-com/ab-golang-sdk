
# List Components Price Points Response

## Structure

`ListComponentsPricePointsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.ComponentPricePoint`](../../doc/models/component-price-point.md) | Required | - |

## Example (as JSON)

```json
{
  "price_points": [
    {
      "id": 40,
      "type": "default",
      "default": false,
      "name": "name2",
      "pricing_scheme": "per_unit"
    }
  ]
}
```

