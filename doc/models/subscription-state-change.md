
# Subscription State Change

## Structure

`SubscriptionStateChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousSubscriptionState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewSubscriptionState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionStateChange := models.SubscriptionStateChange{
        PreviousSubscriptionState: "previous_subscription_state8",
        NewSubscriptionState:      "new_subscription_state2",
    }

}
```

