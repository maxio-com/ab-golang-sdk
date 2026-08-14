
# Multi Invoice Payment

## Structure

`MultiInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `*int` | Optional | The numeric ID of the transaction. |
| `TotalAmount` | `*string` | Optional | Dollar amount of the sum of the paid invoices. |
| `CurrencyCode` | `*string` | Optional | The ISO 4217 currency code (3 character string) representing the currency of invoice transaction. |
| `Applications` | [`[]models.InvoicePaymentApplication`](../../doc/models/invoice-payment-application.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    multiInvoicePayment := models.MultiInvoicePayment{
        TransactionId:        models.ToPointer(246),
        TotalAmount:          models.ToPointer("total_amount0"),
        CurrencyCode:         models.ToPointer("currency_code0"),
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
        },
    }

}
```

