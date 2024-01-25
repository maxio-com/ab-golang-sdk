
# Create on Off Component

## Structure

`CreateOnOffComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OnOffComponent` | [`models.OnOffComponent`](../../doc/models/on-off-component.md) | Required | - |

## Example (as JSON)

```json
{
  "on_off_component": {
    "name": "name6",
    "description": "description6",
    "handle": "handle2",
    "taxable": false,
    "prices": [
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
    ],
    "upgrade_charge": "full"
  }
}
```

