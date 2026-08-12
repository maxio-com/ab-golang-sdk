
# Subscription MRR Response

## Structure

`SubscriptionMRRResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionsMrr` | [`[]models.SubscriptionMRR`](../../doc/models/subscription-mrr.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionMRRResponse := models.SubscriptionMRRResponse{
        SubscriptionsMrr:     []models.SubscriptionMRR{
            models.SubscriptionMRR{
                SubscriptionId:       0,
                MrrAmountInCents:     int64(0),
                Breakouts:            models.ToPointer(models.SubscriptionMRRBreakout{
                    PlanAmountInCents:    int64(0),
                    UsageAmountInCents:   int64(0),
                }),
            },
        },
    }

}
```

