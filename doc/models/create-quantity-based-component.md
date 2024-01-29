
# Create Quantity Based Component

## Structure

`CreateQuantityBasedComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `QuantityBasedComponent` | [`models.QuantityBasedComponent`](../../doc/models/quantity-based-component.md) | Required | - |

## Example (as JSON)

```json
{
  "quantity_based_component": {
    "name": "name0",
    "unit_name": "unit_name2",
    "description": "description0",
    "handle": "handle6",
    "taxable": false,
    "pricing_scheme": "stairstep",
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
      },
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
    "upgrade_charge": "prorated"
  }
}
```
