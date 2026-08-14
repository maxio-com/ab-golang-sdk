
# Remove Payment Event Data

Example schema for an `remove_payment` event

## Structure

`RemovePaymentEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `int` | Required | Transaction ID of the original payment that was removed |
| `Memo` | `string` | Required | Memo of the original payment |
| `OriginalAmount` | `*string` | Optional | Full amount of the original payment |
| `AppliedAmount` | `string` | Required | Applied amount of the original payment |
| `TransactionTime` | `time.Time` | Required | Transaction time of the original payment, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |
| `PaymentMethod` | [`models.InvoiceEventPayment`](../../doc/models/containers/invoice-event-payment.md) | Required | A nested data structure detailing the method of payment |
| `Prepayment` | `bool` | Required | The flag that shows whether the original payment was a prepayment or not |

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
    removePaymentEventData := models.RemovePaymentEventData{
        TransactionId:        20,
        Memo:                 "memo4",
        OriginalAmount:       models.ToPointer("original_amount4"),
        AppliedAmount:        "applied_amount8",
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        PaymentMethod:        models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay{
            Type:                 models.InvoiceEventPaymentMethod_APPLEPAY,
        }),
        Prepayment:           false,
    }

}
```

