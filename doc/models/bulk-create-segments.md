
# Bulk Create Segments

## Structure

`BulkCreateSegments`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segments` | [`[]models.CreateSegment`](../../doc/models/create-segment.md) | Optional | **Constraints**: *Maximum Items*: `2000` |

## Example (as JSON)

```json
{
  "segments": [
    {
      "segment_property_1_value": "String3",
      "segment_property_2_value": "String5",
      "segment_property_3_value": "String3",
      "segment_property_4_value": "String7",
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

