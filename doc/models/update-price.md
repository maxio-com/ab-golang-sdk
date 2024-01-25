
# Update Price

## Structure

`UpdatePrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `EndingQuantity` | `*interface{}` | Optional | - |
| `UnitPrice` | `*interface{}` | Optional | The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065 |
| `Destroy` | `*bool` | Optional | - |
| `StartingQuantity` | `*interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": 18,
  "ending_quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "unit_price": {
    "key1": "val1",
    "key2": "val2"
  },
  "_destroy": false,
  "starting_quantity": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

