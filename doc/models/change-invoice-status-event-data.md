
# Change Invoice Status Event Data

Example schema for an `change_invoice_status` event

## Structure

`ChangeInvoiceStatusEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GatewayTransId` | `*string` | Optional | Identifier for the transaction within the payment gateway. |
| `Amount` | `*string` | Optional | The monetary value associated with the linked payment, expressed in dollars. |
| `FromStatus` | [`models.InvoiceStatus`](../../doc/models/invoice-status.md) | Required | The status of the invoice before any changes occurred. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more. |
| `ToStatus` | [`models.InvoiceStatus`](../../doc/models/invoice-status.md) | Required | The updated status of the invoice after changes have been made. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more. |
| `ConsolidationLevel` | [`*models.InvoiceConsolidationLevel`](../../doc/models/invoice-consolidation-level.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    changeInvoiceStatusEventData := models.ChangeInvoiceStatusEventData{
        GatewayTransId:       models.ToPointer("gateway_trans_id4"),
        Amount:               models.ToPointer("amount6"),
        FromStatus:           models.InvoiceStatus_CANCELED,
        ToStatus:             models.InvoiceStatus_PROCESSING,
        ConsolidationLevel:   models.ToPointer(models.InvoiceConsolidationLevel_PARENT),
    }

}
```

