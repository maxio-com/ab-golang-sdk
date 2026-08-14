
# Subscription Group Members Array Error

## Structure

`SubscriptionGroupMembersArrayError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | `[]string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupMembersArrayError := models.SubscriptionGroupMembersArrayError{
        Members:              []string{
            "members6",
        },
    }

}
```

