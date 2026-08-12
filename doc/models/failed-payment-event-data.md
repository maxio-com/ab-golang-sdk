
# Failed Payment Event Data

Example schema for an `failed_payment` event

## Structure

`FailedPaymentEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmountInCents` | `int` | Required | The monetary value of the payment, expressed in cents. |
| `AppliedAmount` | `int` | Required | The monetary value of the payment, expressed in dollars. |
| `Memo` | `models.Optional[string]` | Optional | The memo passed when the payment was created. |
| `PaymentMethod` | [`models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Required | - |
| `TransactionId` | `int` | Required | The transaction ID of the failed payment. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    failedPaymentEventData := models.FailedPaymentEventData{
        AmountInCents:        34,
        AppliedAmount:        196,
        Memo:                 models.NewOptional(models.ToPointer("memo0")),
        PaymentMethod:        models.InvoicePaymentMethodType_ACH,
        TransactionId:        76,
    }

}
```

