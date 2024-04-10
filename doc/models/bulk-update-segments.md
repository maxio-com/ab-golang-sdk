
# Bulk Update Segments

## Structure

`BulkUpdateSegments`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | [`[]models.BulkUpdateSegmentsItem`](../../doc/models/bulk-update-segments-item.md) | Optional | **Constraints**: *Maximum Items*: `1000` |

## Example (as JSON)

```json
{
  "segments": [
    {
      "id": 50,
      "pricing_scheme": "stairstep",
      "prices": [
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        }
      ]
    },
    {
      "id": 50,
      "pricing_scheme": "stairstep",
      "prices": [
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": "String3"
        }
      ]
    }
  ]
}
```

