
# Create Subscription Component

## Structure

`CreateSubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*interface{}` | Optional | - |
| `Enabled` | `*bool` | Optional | Used for on/off components only. |
| `UnitBalance` | `*int` | Optional | Used for metered and events based components. |
| `AllocatedQuantity` | `*int` | Optional | Used for quantity based components. |
| `Quantity` | `*int` | Optional | Deprecated. Use `allocated_quantity` instead. |
| `PricePointId` | `*interface{}` | Optional | - |
| `CustomPrice` | [`*models.ComponentCustomPrice`](component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example (as JSON)

```json
{
  "component_id": {
    "key1": "val1",
    "key2": "val2"
  },
  "enabled": false,
  "unit_balance": 144,
  "allocated_quantity": 180,
  "quantity": 188
}
```

