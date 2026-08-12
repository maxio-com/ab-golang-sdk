
# Record Payment Response

## Structure

`RecordPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaidInvoices` | [`[]models.PaidInvoice`](../../doc/models/paid-invoice.md) | Optional | - |
| `Prepayment` | [`models.Optional[models.InvoicePrePayment]`](../../doc/models/invoice-pre-payment.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    recordPaymentResponse := models.RecordPaymentResponse{
        PaidInvoices:         []models.PaidInvoice{
            models.PaidInvoice{
                InvoiceId:            models.ToPointer("invoice_id8"),
                Status:               models.ToPointer(models.InvoiceStatus_DRAFT),
                DueAmount:            models.ToPointer("due_amount0"),
                PaidAmount:           models.ToPointer("paid_amount0"),
            },
        },
        Prepayment:           models.NewOptional(models.ToPointer(models.InvoicePrePayment{
            SubscriptionId:       models.ToPointer(148),
            AmountInCents:        models.ToPointer(int64(124)),
            EndingBalanceInCents: models.ToPointer(int64(164)),
        })),
    }

}
```

