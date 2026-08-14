
# List Subscription Groups Meta

## Structure

`ListSubscriptionGroupsMeta`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentPage` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupsMeta := models.ListSubscriptionGroupsMeta{
        CurrentPage:          models.ToPointer(104),
        TotalCount:           models.ToPointer(128),
    }

}
```

