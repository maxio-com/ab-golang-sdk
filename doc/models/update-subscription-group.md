
# Update Subscription Group

## Structure

`UpdateSubscriptionGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MemberIds` | `[]int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSubscriptionGroup := models.UpdateSubscriptionGroup{
        MemberIds:            []int{
            248,
            249,
            250,
        },
    }

}
```

