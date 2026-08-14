
# Reactivate Subscription Request

## Structure

`ReactivateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CalendarBilling` | [`*models.ReactivationBilling`](../../doc/models/reactivation-billing.md) | Optional | These values are only applicable to subscriptions using calendar billing. |
| `IncludeTrial` | `*bool` | Optional | If `true` is sent, the reactivated Subscription will include a trial if one is available. If `false` is sent, the trial period will be ignored. |
| `PreserveBalance` | `*bool` | Optional | If `true` is passed, the existing subscription balance will NOT be cleared/reset before adding the additional reactivation charges. |
| `CouponCode` | `*string` | Optional | The coupon code to be applied during reactivation. |
| `UseCreditsAndPrepayments` | `*bool` | Optional | If true is sent, Advanced Billing will use service credits and prepayments upon reactivation. If false is sent, the service credits and prepayments will be ignored. |
| `Resume` | [`*models.ReactivateSubscriptionRequestResume`](../../doc/models/containers/reactivate-subscription-request-resume.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reactivateSubscriptionRequest := models.ReactivateSubscriptionRequest{
        CalendarBilling:          models.ToPointer(models.ReactivationBilling{
            ReactivationCharge:   models.ToPointer(models.ReactivationCharge_PRORATED),
        }),
        IncludeTrial:             models.ToPointer(false),
        PreserveBalance:          models.ToPointer(false),
        CouponCode:               models.ToPointer("coupon_code2"),
        UseCreditsAndPrepayments: models.ToPointer(false),
    }

}
```

