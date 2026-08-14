
# Refund Consolidated Invoice

Refund consolidated invoice.

## Structure

`RefundConsolidatedInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Memo` | `string` | Required | A description for the refund |
| `PaymentId` | `int` | Required | The ID of the payment to be refunded |
| `SegmentUids` | [`models.RefundConsolidatedInvoiceSegmentUids`](../../doc/models/containers/refund-consolidated-invoice-segment-uids.md) | Required | This is a container for one-of cases. |
| `External` | `*bool` | Optional | Flag that marks refund as external (no money is returned to the customer). Defaults to `false`. |
| `ApplyCredit` | `*bool` | Optional | If set to true, creates credit and applies it to an invoice. Defaults to `false`. |
| `Amount` | `*string` | Optional | The amount of payment to be refunded in decimal format. Example: "10.50". This will default to the full amount of the payment if not provided. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundConsolidatedInvoice := models.RefundConsolidatedInvoice{
        Memo:                 "memo0",
        PaymentId:            138,
        SegmentUids:          models.RefundConsolidatedInvoiceSegmentUidsContainer.FromArrayOfString([]string{
            "String0",
            "String1",
        }),
        External:             models.ToPointer(false),
        ApplyCredit:          models.ToPointer(false),
        Amount:               models.ToPointer("amount8"),
    }

}
```

