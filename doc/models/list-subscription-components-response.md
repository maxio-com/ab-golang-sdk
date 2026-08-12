
# List Subscription Components Response

## Structure

`ListSubscriptionComponentsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionsComponents` | [`[]models.SubscriptionComponent`](../../doc/models/subscription-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionComponentsResponse := models.ListSubscriptionComponentsResponse{
        SubscriptionsComponents: []models.SubscriptionComponent{
            models.SubscriptionComponent{
                Id:                        models.ToPointer(138),
                Name:                      models.ToPointer("name2"),
                Kind:                      models.ToPointer(models.ComponentKind_METEREDCOMPONENT),
                UnitName:                  models.ToPointer("unit_name4"),
                Enabled:                   models.ToPointer(false),
            },
        },
    }

}
```

