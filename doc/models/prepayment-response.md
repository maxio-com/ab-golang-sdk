
# Prepayment Response

## Structure

`PrepaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.Prepayment`](../../doc/models/prepayment.md) | Required | - |

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
    prepaymentResponse := models.PrepaymentResponse{
        Prepayment:           models.Prepayment{
            Id:                     38,
            SubscriptionId:         148,
            AmountInCents:          int64(124),
            RemainingAmountInCents: int64(182),
            RefundedAmountInCents:  models.ToPointer(int64(132)),
            Details:                models.ToPointer("details8"),
            External:               false,
            Memo:                   "memo2",
            PaymentType:            models.ToPointer(models.PrepaymentMethod_CREDITCARD),
            CreatedAt:              parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        },
    }

}
```

