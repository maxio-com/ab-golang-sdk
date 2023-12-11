
# Price

## Structure

`Price`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `interface{}` | Required | - |
| `EndingQuantity` | `Optional[interface{}]` | Optional | - |
| `UnitPrice` | `interface{}` | Required | The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065 |

## Example (as JSON)

```json
{
  "starting_quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "ending_quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "unit_price": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

