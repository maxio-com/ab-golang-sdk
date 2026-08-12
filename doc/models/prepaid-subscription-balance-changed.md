
# Prepaid Subscription Balance Changed

## Structure

`PrepaidSubscriptionBalanceChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `string` | Required | - |
| `CurrentAccountBalanceInCents` | `int64` | Required | - |
| `PrepaymentAccountBalanceInCents` | `int64` | Required | - |
| `CurrentUsageAmountInCents` | `int64` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidSubscriptionBalanceChanged := models.PrepaidSubscriptionBalanceChanged{
        Reason:                          "reason6",
        CurrentAccountBalanceInCents:    int64(194),
        PrepaymentAccountBalanceInCents: int64(100),
        CurrentUsageAmountInCents:       int64(186),
    }

}
```

