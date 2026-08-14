
# Update Subscription Group Request

## Structure

`UpdateSubscriptionGroupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.UpdateSubscriptionGroup`](../../doc/models/update-subscription-group.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionGroupRequest := models.UpdateSubscriptionGroupRequest{
        SubscriptionGroup:    models.UpdateSubscriptionGroup{
            MemberIds:            []int{
                164,
                165,
            },
        },
    }

}
```

