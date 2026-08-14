
# Create Invoice Payment Request

## Structure

`CreateInvoicePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.CreateInvoicePayment`](../../doc/models/create-invoice-payment.md) | Required | - |
| `Type` | [`*models.InvoicePaymentType`](../../doc/models/invoice-payment-type.md) | Optional | The type of payment to be applied to an Invoice. Defaults to external. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createInvoicePaymentRequest := models.CreateInvoicePaymentRequest{
        Payment:              models.CreateInvoicePayment{
            Amount:               models.ToPointer(models.CreateInvoicePaymentAmountContainer.FromString("String9")),
            Memo:                 models.ToPointer("memo0"),
            Method:               models.ToPointer(models.InvoicePaymentMethodType_ACH),
            Details:              models.ToPointer("details6"),
            PaymentProfileId:     models.ToPointer(42),
        },
        Type:                 models.ToPointer(models.InvoicePaymentType_EXTERNAL),
    }

}
```

