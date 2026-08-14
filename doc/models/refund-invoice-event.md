
# Refund Invoice Event

## Structure

`RefundInvoiceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"refund_invoice"` |
| `EventData` | [`models.RefundInvoiceEventData`](../../doc/models/refund-invoice-event-data.md) | Required | Example schema for an `refund_invoice` event |

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
    refundInvoiceEvent := models.RefundInvoiceEvent{
        Id:                   int64(132),
        Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        Invoice:              models.Invoice{
            Id:                         models.ToPointer(int64(166)),
            Uid:                        models.ToPointer("uid6"),
            SiteId:                     models.ToPointer(92),
            CustomerId:                 models.ToPointer(204),
            SubscriptionId:             models.ToPointer(20),
            IssueDate:                  models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
            DueDate:                    models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
            PaidDate:                   models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) }))),
            PublicUrlExpiresOn:         models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-21", func(err error) { log.Fatalln(err) })),
        },
        EventType:            models.InvoiceEventType_REFUNDINVOICE,
        EventData:            models.RefundInvoiceEventData{
            ApplyCredit:          false,
            ConsolidationLevel:   models.ToPointer(models.InvoiceConsolidationLevel_CHILD),
            CreditNoteAttributes: models.CreditNote{
                Uid:                  models.ToPointer("uid2"),
                SiteId:               models.ToPointer(72),
                CustomerId:           models.ToPointer(184),
                SubscriptionId:       models.ToPointer(0),
                Number:               models.ToPointer("number0"),
            },
            Memo:                 models.ToPointer("memo0"),
            OriginalAmount:       models.ToPointer("original_amount0"),
            PaymentId:            204,
            RefundAmount:         "refund_amount8",
            RefundId:             248,
            TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        },
    }

}
```

