
# Delete Subscription Group Response

## Structure

`DeleteSubscriptionGroupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Deleted` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    deleteSubscriptionGroupResponse := models.DeleteSubscriptionGroupResponse{
        Uid:                  models.ToPointer("uid0"),
        Deleted:              models.ToPointer(false),
    }

}
```

