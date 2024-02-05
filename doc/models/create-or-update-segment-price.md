
# Create or Update Segment Price

## Structure

`CreateOrUpdateSegmentPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `*int` | Optional | - |
| `UnitPrice` | `interface{}` | Required | The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065 |

## Example (as JSON)

```json
{
  "starting_quantity": 78,
  "ending_quantity": 52,
  "unit_price": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

