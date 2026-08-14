
# Refund Invoice Request

## Structure

`RefundInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`models.RefundInvoiceRequestRefund`](../../doc/models/containers/refund-invoice-request-refund.md) | Required | This is a container for any-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundInvoiceRequest := models.RefundInvoiceRequest{
        Refund:               models.RefundInvoiceRequestRefundContainer.FromRefundInvoice(models.RefundInvoice{
            Amount:               "amount8",
            Memo:                 "memo0",
            PaymentId:            0,
            External:             models.ToPointer(false),
            ApplyCredit:          models.ToPointer(false),
            VoidInvoice:          models.ToPointer(false),
        }),
    }

}
```

