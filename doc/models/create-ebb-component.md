
# Create EBB Component

## Structure

`CreateEBBComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EventBasedComponent` | [`models.EBBComponent`](../../doc/models/ebb-component.md) | Required | - |

## Example (as JSON)

```json
{
  "event_based_component": {
    "name": "name8",
    "unit_name": "unit_name0",
    "description": "description8",
    "handle": "handle4",
    "taxable": false,
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ],
    "upgrade_charge": "full",
    "event_based_billing_metric_id": 68
  }
}
```

