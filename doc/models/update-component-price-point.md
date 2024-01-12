
# Update Component Price Point

## Structure

`UpdateComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `Prices` | [`[]models.UpdatePrice`](../../doc/models/update-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "interval": 216,
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
    }
  ]
}
```

