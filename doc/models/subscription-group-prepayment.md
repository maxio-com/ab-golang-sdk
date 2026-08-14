
# Subscription Group Prepayment

## Structure

`SubscriptionGroupPrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `int` | Required | - |
| `Details` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `Method` | [`models.SubscriptionGroupPrepaymentMethod`](../../doc/models/subscription-group-prepayment-method.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupPrepayment := models.SubscriptionGroupPrepayment{
        Amount:               12,
        Details:              "details4",
        Memo:                 "memo8",
        Method:               models.SubscriptionGroupPrepaymentMethod_MONEYORDER,
    }

}
```

