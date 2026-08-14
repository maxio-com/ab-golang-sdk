
# Subscription Group Prepayment Request

## Structure

`SubscriptionGroupPrepaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.SubscriptionGroupPrepayment`](../../doc/models/subscription-group-prepayment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupPrepaymentRequest := models.SubscriptionGroupPrepaymentRequest{
        Prepayment:           models.SubscriptionGroupPrepayment{
            Amount:               136,
            Details:              "details8",
            Memo:                 "memo2",
            Method:               models.SubscriptionGroupPrepaymentMethod_PAYPALACCOUNT,
        },
    }

}
```

