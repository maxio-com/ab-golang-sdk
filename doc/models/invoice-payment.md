
# Invoice Payment

## Structure

`InvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionTime` | `*time.Time` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |
| `PaymentMethod` | [`*models.InvoicePaymentMethod`](../../doc/models/invoice-payment-method.md) | Optional | - |
| `TransactionId` | `*int` | Optional | - |
| `Prepayment` | `*bool` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `GatewayUsed` | `*string` | Optional | - |
| `GatewayTransactionId` | `models.Optional[string]` | Optional | The transaction ID for the payment as returned from the payment gateway |
| `ReceivedOn` | `models.Optional[time.Time]` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. Applicable only to<br>`external` payments. |
| `Uid` | `*string` | Optional | - |

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
    invoicePayment := models.InvoicePayment{
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Memo:                 models.ToPointer("memo6"),
        OriginalAmount:       models.ToPointer("original_amount6"),
        AppliedAmount:        models.ToPointer("applied_amount6"),
        PaymentMethod:        models.ToPointer(models.InvoicePaymentMethod{
            Details:              models.ToPointer("details0"),
            Kind:                 models.ToPointer("kind8"),
            Memo:                 models.ToPointer("memo4"),
            Type:                 models.ToPointer("type0"),
            CardBrand:            models.ToPointer("card_brand6"),
        }),
    }

}
```

