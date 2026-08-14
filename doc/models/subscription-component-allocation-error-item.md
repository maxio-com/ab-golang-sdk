
# Subscription Component Allocation Error Item

## Structure

`SubscriptionComponentAllocationErrorItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Kind` | `*string` | Optional | - |
| `Message` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionComponentAllocationErrorItem := models.SubscriptionComponentAllocationErrorItem{
        Kind:                 models.ToPointer("kind6"),
        Message:              models.ToPointer("message8"),
    }

}
```

