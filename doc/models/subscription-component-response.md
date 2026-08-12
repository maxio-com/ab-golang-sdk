
# Subscription Component Response

## Structure

`SubscriptionComponentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Component` | [`*models.SubscriptionComponent`](../../doc/models/subscription-component.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionComponentResponse := models.SubscriptionComponentResponse{
        Component:            models.ToPointer(models.SubscriptionComponent{
            Id:                        models.ToPointer(80),
            Name:                      models.ToPointer("name8"),
            Kind:                      models.ToPointer(models.ComponentKind_QUANTITYBASEDCOMPONENT),
            UnitName:                  models.ToPointer("unit_name0"),
            Enabled:                   models.ToPointer(false),
        }),
    }

}
```

