
# Subscription MRR

## Structure

`SubscriptionMRR`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `int` | Required | - |
| `MrrAmountInCents` | `int64` | Required | - |
| `Breakouts` | [`*models.SubscriptionMRRBreakout`](../../doc/models/subscription-mrr-breakout.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMRR := models.SubscriptionMRR{
        SubscriptionId:       192,
        MrrAmountInCents:     int64(210),
        Breakouts:            models.ToPointer(models.SubscriptionMRRBreakout{
            PlanAmountInCents:    int64(254),
            UsageAmountInCents:   int64(106),
        }),
    }

}
```

