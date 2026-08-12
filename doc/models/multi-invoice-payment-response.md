
# Multi Invoice Payment Response

## Structure

`MultiInvoicePaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.MultiInvoicePayment`](../../doc/models/multi-invoice-payment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    multiInvoicePaymentResponse := models.MultiInvoicePaymentResponse{
        Payment:              models.MultiInvoicePayment{
            TransactionId:        models.ToPointer(224),
            TotalAmount:          models.ToPointer("total_amount2"),
            CurrencyCode:         models.ToPointer("currency_code2"),
            Applications:         []models.InvoicePaymentApplication{
                models.InvoicePaymentApplication{
                    InvoiceUid:           models.ToPointer("invoice_uid8"),
                    ApplicationUid:       models.ToPointer("application_uid8"),
                    AppliedAmount:        models.ToPointer("applied_amount0"),
                },
                models.InvoicePaymentApplication{
                    InvoiceUid:           models.ToPointer("invoice_uid8"),
                    ApplicationUid:       models.ToPointer("application_uid8"),
                    AppliedAmount:        models.ToPointer("applied_amount0"),
                },
                models.InvoicePaymentApplication{
                    InvoiceUid:           models.ToPointer("invoice_uid8"),
                    ApplicationUid:       models.ToPointer("application_uid8"),
                    AppliedAmount:        models.ToPointer("applied_amount0"),
                },
            },
        },
    }

}
```

