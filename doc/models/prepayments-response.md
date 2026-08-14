
# Prepayments Response

## Structure

`PrepaymentsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`[]models.Prepayment`](../../doc/models/prepayment.md) | Optional | **Constraints**: *Unique Items Required* |

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
    prepaymentsResponse := models.PrepaymentsResponse{
        Prepayments:          []models.Prepayment{
            models.Prepayment{
                Id:                     76,
                SubscriptionId:         186,
                AmountInCents:          int64(94),
                RemainingAmountInCents: int64(220),
                RefundedAmountInCents:  models.ToPointer(int64(170)),
                Details:                models.ToPointer("details6"),
                External:               false,
                Memo:                   "memo0",
                PaymentType:            models.ToPointer(models.PrepaymentMethod_CASH),
                CreatedAt:              parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
            },
        },
    }

}
```

