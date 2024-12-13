
# Create Metered Component

## Structure

`CreateMeteredComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MeteredComponent` | [`models.MeteredComponent`](../../doc/models/metered-component.md) | Required | - |

## Example (as JSON)

```json
{
  "metered_component": {
    "name": "name0",
    "unit_name": "unit_name2",
    "description": "description0",
    "handle": "handle6",
    "taxable": false,
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      },
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      },
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ],
    "price_points": [
      {
        "name": "name2",
        "handle": "handle8",
        "pricing_scheme": "per_unit",
        "interval": 92,
        "interval_unit": "day"
      },
      {
        "name": "name2",
        "handle": "handle8",
        "pricing_scheme": "per_unit",
        "interval": 92,
        "interval_unit": "day"
      },
      {
        "name": "name2",
        "handle": "handle8",
        "pricing_scheme": "per_unit",
        "interval": 92,
        "interval_unit": "day"
      }
    ]
  }
}
```

