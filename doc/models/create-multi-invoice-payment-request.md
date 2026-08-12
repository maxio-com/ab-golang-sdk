
# Create Multi Invoice Payment Request

## Structure

`CreateMultiInvoicePaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.CreateMultiInvoicePayment`](../../doc/models/create-multi-invoice-payment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMultiInvoicePaymentRequest := models.CreateMultiInvoicePaymentRequest{
        Payment:              models.CreateMultiInvoicePayment{
            Memo:                 models.ToPointer("memo0"),
            Details:              models.ToPointer("details6"),
            Method:               models.ToPointer(models.InvoicePaymentMethodType_ACH),
            Amount:               models.CreateMultiInvoicePaymentAmountContainer.FromString("String9"),
            ReceivedOn:           models.ToPointer("received_on8"),
            Applications:         []models.CreateInvoicePaymentApplication{
                models.CreateInvoicePaymentApplication{
                    InvoiceUid:           "invoice_uid8",
                    Amount:               "amount0",
                },
            },
        },
    }

}
```

