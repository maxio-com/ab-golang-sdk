
# Subscription Included Coupon

## Structure

`SubscriptionIncludedCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | - |
| `UseCount` | `*int` | Optional | - |
| `UsesAllowed` | `*int` | Optional | - |
| `ExpiresAt` | `models.Optional[string]` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `AmountInCents` | `models.Optional[int64]` | Optional | **Constraints**: `>= 0` |
| `Percentage` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionIncludedCoupon := models.SubscriptionIncludedCoupon{
        Code:                 models.ToPointer("\"ABCD_10\""),
        UseCount:             models.ToPointer(2),
        UsesAllowed:          models.ToPointer(10),
        ExpiresAt:            models.NewOptional(models.ToPointer("\"2023-07-13T05:18:58-04:00\"")),
        Recurring:            models.ToPointer(false),
        AmountInCents:        models.NewOptional(models.ToPointer(int64(1000))),
        Percentage:           models.NewOptional(models.ToPointer("\"15.0\"")),
    }

}
```

