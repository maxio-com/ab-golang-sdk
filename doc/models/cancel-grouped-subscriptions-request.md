
# Cancel Grouped Subscriptions Request

## Structure

`CancelGroupedSubscriptionsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargeUnbilledUsage` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cancelGroupedSubscriptionsRequest := models.CancelGroupedSubscriptionsRequest{
        ChargeUnbilledUsage:  models.ToPointer(false),
    }

}
```

