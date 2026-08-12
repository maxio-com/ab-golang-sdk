
# Create Invoice Payment

## Structure

`CreateInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | [`*models.CreateInvoicePaymentAmount`](../../doc/models/containers/create-invoice-payment-amount.md) | Optional | This is a container for one-of cases. |
| `Memo` | `*string` | Optional | A description to be attached to the payment. Applicable only to `external` payments. |
| `Method` | [`*models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Optional | The type of payment method used. Defaults to other. |
| `Details` | `*string` | Optional | Additional information related to the payment method (eg. Check #). Applicable only to `external` payments. |
| `PaymentProfileId` | `*int` | Optional | The ID of the payment profile to be used for the payment. |
| `ReceivedOn` | `*time.Time` | Optional | Date reflecting when the payment was received from a customer. Must be in the past. Applicable only to<br>`external` payments. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createInvoicePayment := models.CreateInvoicePayment{
        Amount:               models.ToPointer(models.CreateInvoicePaymentAmountContainer.FromString("String7")),
        Memo:                 models.ToPointer("memo8"),
        Method:               models.ToPointer(models.InvoicePaymentMethodType_ACH),
        Details:              models.ToPointer("details4"),
        PaymentProfileId:     models.ToPointer(30),
    }

}
```

