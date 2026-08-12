
# Create Debit Note Event

## Structure

`CreateDebitNoteEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"create_debit_note"` |
| `EventData` | [`models.DebitNote`](../../doc/models/debit-note.md) | Required | Example schema for an `create_debit_note` event |

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
    createDebitNoteEvent := models.CreateDebitNoteEvent{
        Id:                   int64(176),
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
        EventType:            models.InvoiceEventType_CREATEDEBITNOTE,
        EventData:            models.DebitNote{
            Uid:                    models.ToPointer("uid6"),
            SiteId:                 models.ToPointer(132),
            CustomerId:             models.ToPointer(244),
            SubscriptionId:         models.ToPointer(60),
            Number:                 models.ToPointer(64),
        },
    }

}
```

