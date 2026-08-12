
# Prepayment Account Balance Changed

## Structure

`PrepaymentAccountBalanceChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `string` | Required | - |
| `PrepaymentAccountBalanceInCents` | `int64` | Required | - |
| `PrepaymentBalanceChangeInCents` | `int64` | Required | - |
| `CurrencyCode` | `string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaymentAccountBalanceChanged := models.PrepaymentAccountBalanceChanged{
        Reason:                          "reason8",
        PrepaymentAccountBalanceInCents: int64(134),
        PrepaymentBalanceChangeInCents:  int64(158),
        CurrencyCode:                    "currency_code8",
    }

}
```

