
# Reactivate Subscription Group Request

## Structure

`ReactivateSubscriptionGroupRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Resume` | `*bool` | Optional | - |
| `ResumeMembers` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reactivateSubscriptionGroupRequest := models.ReactivateSubscriptionGroupRequest{
        Resume:               models.ToPointer(false),
        ResumeMembers:        models.ToPointer(false),
    }

}
```

