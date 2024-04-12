
# Update Segment Request

## Structure

`UpdateSegmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Segment` | [`models.UpdateSegment`](../../doc/models/update-segment.md) | Required | - |

## Example (as JSON)

```json
{
  "segment": {
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
}
```

