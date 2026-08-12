
# Subscription MRR Breakout

## Structure

`SubscriptionMRRBreakout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PlanAmountInCents` | `int64` | Required | - |
| `UsageAmountInCents` | `int64` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMRRBreakout := models.SubscriptionMRRBreakout{
        PlanAmountInCents:    int64(248),
        UsageAmountInCents:   int64(100),
    }

}
```

