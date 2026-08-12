
# Apply Payment Event Data

Example schema for an `apply_payment` event

## Structure

`ApplyPaymentEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConsolidationLevel` | [`models.InvoiceConsolidationLevel`](../../doc/models/invoice-consolidation-level.md) | Required | - |
| `Memo` | `string` | Required | The payment memo |
| `OriginalAmount` | `string` | Required | The full, original amount of the payment transaction as a string in full units. Incoming payments can be split amongst several invoices, which will result in a `applied_amount` less than the `original_amount`. Example: A $100.99 payment, of which $40.11 is applied to this invoice, will have an `original_amount` of `"100.99"`. |
| `AppliedAmount` | `string` | Required | The amount of the payment applied to this invoice. Incoming payments can be split amongst several invoices, which will result in a `applied_amount` less than the `original_amount`. Example: A $100.99 payment, of which $40.11 is applied to this invoice, will have an `applied_amount` of `"40.11"`. |
| `TransactionTime` | `time.Time` | Required | The time the payment was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |
| `PaymentMethod` | [`models.InvoiceEventPayment`](../../doc/models/containers/invoice-event-payment.md) | Required | A nested data structure detailing the method of payment |
| `TransactionId` | `*int` | Optional | The Chargify id of the original payment |
| `ParentInvoiceNumber` | `models.Optional[int]` | Optional | - |
| `RemainingPrepaymentAmount` | `models.Optional[string]` | Optional | - |
| `Prepayment` | `*bool` | Optional | - |
| `External` | `*bool` | Optional | - |

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
    applyPaymentEventData := models.ApplyPaymentEventData{
        ConsolidationLevel:        models.InvoiceConsolidationLevel_CHILD,
        Memo:                      "memo0",
        OriginalAmount:            "original_amount0",
        AppliedAmount:             "applied_amount2",
        TransactionTime:           parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        PaymentMethod:             models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay{
            Type:                 models.InvoiceEventPaymentMethod_APPLEPAY,
        }),
        TransactionId:             models.ToPointer(142),
        ParentInvoiceNumber:       models.NewOptional(models.ToPointer(228)),
        RemainingPrepaymentAmount: models.NewOptional(models.ToPointer("remaining_prepayment_amount4")),
        Prepayment:                models.ToPointer(false),
        External:                  models.ToPointer(false),
    }

}
```

