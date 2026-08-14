
# Paid Invoice

## Structure

`PaidInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceId` | `*string` | Optional | The uid of the paid invoice |
| `Status` | [`*models.InvoiceStatus`](../../doc/models/invoice-status.md) | Optional | The current status of the invoice. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more. |
| `DueAmount` | `*string` | Optional | The remaining due amount on the invoice |
| `PaidAmount` | `*string` | Optional | The total amount paid on this invoice (including any prior payments) |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paidInvoice := models.PaidInvoice{
        InvoiceId:            models.ToPointer("invoice_id2"),
        Status:               models.ToPointer(models.InvoiceStatus_CANCELED),
        DueAmount:            models.ToPointer("due_amount4"),
        PaidAmount:           models.ToPointer("paid_amount6"),
    }

}
```

