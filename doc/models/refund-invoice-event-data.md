
# Refund Invoice Event Data

Example schema for an `refund_invoice` event

## Structure

`RefundInvoiceEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApplyCredit` | `bool` | Required | If true, credit was created and applied it to the invoice. |
| `ConsolidationLevel` | [`*models.InvoiceConsolidationLevel`](../../doc/models/invoice-consolidation-level.md) | Optional | Consolidation level of the invoice, which is applicable to invoice consolidation. It will hold one of the following values:<br><br>* "none": A normal invoice with no consolidation.<br>* "child": An invoice segment which has been combined into a consolidated invoice.<br>* "parent": A consolidated invoice, whose contents are composed of invoice segments.<br><br>"Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.<br><br>See also the [invoice consolidation documentation](https://maxio.zendesk.com/hc/en-us/articles/24252269909389-Invoice-Consolidation). |
| `CreditNoteAttributes` | [`models.CreditNote`](../../doc/models/credit-note.md) | Required | - |
| `Memo` | `*string` | Optional | The refund memo. |
| `OriginalAmount` | `*string` | Optional | The full, original amount of the refund. |
| `PaymentId` | `int` | Required | The ID of the payment transaction to be refunded. |
| `RefundAmount` | `string` | Required | The amount of the refund. |
| `RefundId` | `int` | Required | The ID of the refund transaction. |
| `TransactionTime` | `time.Time` | Required | The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    refundInvoiceEventData := models.RefundInvoiceEventData{
        ApplyCredit:          false,
        ConsolidationLevel:   models.ToPointer(models.InvoiceConsolidationLevel_PARENT),
        CreditNoteAttributes: models.CreditNote{
            Uid:                  models.ToPointer("uid2"),
            SiteId:               models.ToPointer(72),
            CustomerId:           models.ToPointer(184),
            SubscriptionId:       models.ToPointer(0),
            Number:               models.ToPointer("number0"),
        },
        Memo:                 models.ToPointer("memo0"),
        OriginalAmount:       models.ToPointer("original_amount0"),
        PaymentId:            140,
        RefundAmount:         "refund_amount8",
        RefundId:             184,
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    }

}
```

