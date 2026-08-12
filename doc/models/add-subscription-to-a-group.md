
# Add Subscription to a Group

## Structure

`AddSubscriptionToAGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | [`*models.GroupSettings`](../../doc/models/group-settings.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    addSubscriptionToAGroup := models.AddSubscriptionToAGroup{
        Group:                models.ToPointer(models.GroupSettings{
            Target:               models.GroupTarget{
                Type:                 models.GroupTargetType_PARENT,
                Id:                   models.ToPointer(236),
            },
            Billing:              models.ToPointer(models.GroupBilling{
                Accrue:               models.ToPointer(false),
                AlignDate:            models.ToPointer(false),
                Prorate:              models.ToPointer(false),
            }),
        }),
    }

}
```

