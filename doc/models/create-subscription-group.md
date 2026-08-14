
# Create Subscription Group

## Structure

`CreateSubscriptionGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `int` | Required | - |
| `MemberIds` | `[]int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createSubscriptionGroup := models.CreateSubscriptionGroup{
        SubscriptionId:       204,
        MemberIds:            []int{
            48,
        },
    }

}
```

