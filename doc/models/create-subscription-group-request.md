
# Create Subscription Group Request

## Structure

`CreateSubscriptionGroupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.CreateSubscriptionGroup`](../../doc/models/create-subscription-group.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSubscriptionGroupRequest := models.CreateSubscriptionGroupRequest{
        SubscriptionGroup:    models.CreateSubscriptionGroup{
            SubscriptionId:       36,
            MemberIds:            []int{
                164,
                165,
            },
        },
    }

}
```

