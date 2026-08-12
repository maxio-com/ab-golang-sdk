
# Create Multi Invoice Payment

## Structure

`CreateMultiInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `*string` | Optional | A description to be attached to the payment. |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #). |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used. Defaults to other. |
| `Amount` | [`models.CreateMultiInvoicePaymentAmount`](../../doc/models/containers/create-multi-invoice-payment-amount.md) | Required | This is a container for one-of cases. |
| `ReceivedOn` | `*string` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. |
| `Applications` | [`[]models.CreateInvoicePaymentApplication`](../../doc/models/create-invoice-payment-application.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMultiInvoicePayment := models.CreateMultiInvoicePayment{
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
    }

}
```

