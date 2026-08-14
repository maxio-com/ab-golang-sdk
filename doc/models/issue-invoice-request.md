
# Issue Invoice Request

## Structure

`IssueInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OnFailedPayment` | [`*models.FailedPaymentAction`](../../doc/models/failed-payment-action.md) | Optional | Action taken when payment for an invoice fails:<br><br>- `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.<br>- `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.<br>- `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).<br><br>**Default**: `"leave_open_invoice"` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    issueInvoiceRequest := models.IssueInvoiceRequest{
        OnFailedPayment:      models.ToPointer(models.FailedPaymentAction_LEAVEOPENINVOICE),
    }

}
```

