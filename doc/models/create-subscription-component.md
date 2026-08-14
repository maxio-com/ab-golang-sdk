
# Create Subscription Component

## Structure

`CreateSubscriptionComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | [`*models.CreateSubscriptionComponentComponentId`](../../doc/models/containers/create-subscription-component-component-id.md) | Optional | This is a container for one-of cases. |
| `Enabled` | `*bool` | Optional | Used for on/off components only. |
| `UnitBalance` | [`*models.CreateSubscriptionComponentUnitBalance`](../../doc/models/containers/create-subscription-component-unit-balance.md) | Optional | This is a container for one-of cases. |
| `AllocatedQuantity` | [`*models.CreateSubscriptionComponentAllocatedQuantity`](../../doc/models/containers/create-subscription-component-allocated-quantity.md) | Optional | This is a container for one-of cases. |
| `Quantity` | `*int` | Optional | Deprecated. Use `allocated_quantity` instead. |
| `PricePointId` | [`*models.CreateSubscriptionComponentPricePointId`](../../doc/models/containers/create-subscription-component-price-point-id.md) | Optional | This is a container for one-of cases. |
| `CustomPrice` | [`*models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSubscriptionComponent := models.CreateSubscriptionComponent{
        ComponentId:          models.ToPointer(models.CreateSubscriptionComponentComponentIdContainer.FromNumber(210)),
        Enabled:              models.ToPointer(false),
        UnitBalance:          models.ToPointer(models.CreateSubscriptionComponentUnitBalanceContainer.FromNumber(12)),
        AllocatedQuantity:    models.ToPointer(models.CreateSubscriptionComponentAllocatedQuantityContainer.FromNumber(48)),
        Quantity:             models.ToPointer(134),
    }

}
```

