
# Create Segment Request

## Structure

`CreateSegmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`models.CreateSegment`](../../doc/models/create-segment.md) | Required | - |

## Example (as JSON)

```json
{
  "segment": {
    "segment_property_1_value": {
      "key1": "val1",
      "key2": "val2"
    },
    "segment_property_2_value": {
      "key1": "val1",
      "key2": "val2"
    },
    "segment_property_3_value": {
      "key1": "val1",
      "key2": "val2"
    },
    "segment_property_4_value": {
      "key1": "val1",
      "key2": "val2"
    },
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
}
```

