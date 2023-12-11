
# Component Price Points Response

## Structure

`ComponentPricePointsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoints` | [`[]models.ComponentPricePoint`](component-price-point.md) | Optional | - |

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
    },
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

