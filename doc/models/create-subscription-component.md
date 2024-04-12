
# Create Subscription Component

## Structure

`CreateSubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | [`*models.CreateSubscriptionComponentComponentId`](../../doc/models/containers/create-subscription-component-component-id.md) | Optional | This is a container for one-of cases. |
| `Enabled` | `*bool` | Optional | Used for on/off components only. |
| `UnitBalance` | `*int` | Optional | Used for metered and events based components. |
| `AllocatedQuantity` | [`*models.CreateSubscriptionComponentAllocatedQuantity`](../../doc/models/containers/create-subscription-component-allocated-quantity.md) | Optional | This is a container for one-of cases. |
| `Quantity` | `*int` | Optional | Deprecated. Use `allocated_quantity` instead. |
| `PricePointId` | [`*models.CreateSubscriptionComponentPricePointId`](../../doc/models/containers/create-subscription-component-price-point-id.md) | Optional | This is a container for one-of cases. |
| `CustomPrice` | [`*models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example (as JSON)

```json
{
  "component_id": 8,
  "enabled": false,
  "unit_balance": 144,
  "allocated_quantity": 102,
  "quantity": 188
}
```

