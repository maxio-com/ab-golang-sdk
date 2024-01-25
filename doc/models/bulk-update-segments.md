
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
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
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
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
          "unit_price": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "starting_quantity": 64,
          "ending_quantity": 38,
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

