
# Credit Account Balance Changed

## Structure

`CreditAccountBalanceChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `string` | Required | - |
| `ServiceCreditAccountBalanceInCents` | `int64` | Required | - |
| `ServiceCreditBalanceChangeInCents` | `int64` | Required | - |
| `CurrencyCode` | `string` | Required | - |
| `AtTime` | `time.Time` | Required | - |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    creditAccountBalanceChanged := models.CreditAccountBalanceChanged{
        Reason:                             "reason4",
        ServiceCreditAccountBalanceInCents: int64(216),
        ServiceCreditBalanceChangeInCents:  int64(166),
        CurrencyCode:                       "currency_code6",
        AtTime:                             parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    }

}
```

