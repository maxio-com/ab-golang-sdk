
# Apply Credit Note Event

## Structure

`ApplyCreditNoteEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Timestamp` | `time.Time` | Required | - |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |
| `EventType` | [`models.InvoiceEventType`](../../doc/models/invoice-event-type.md) | Required | **Default**: `"apply_credit_note"` |
| `EventData` | [`models.ApplyCreditNoteEventData`](../../doc/models/apply-credit-note-event-data.md) | Required | Example schema for an `apply_credit_note` event |

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
    applyCreditNoteEvent := models.ApplyCreditNoteEvent{
        Id:                   int64(86),
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
        EventType:            models.InvoiceEventType_APPLYCREDITNOTE,
        EventData:            models.ApplyCreditNoteEventData{
            Uid:                  "uid6",
            CreditNoteNumber:     "credit_note_number0",
            CreditNoteUid:        "credit_note_uid0",
            OriginalAmount:       "original_amount0",
            AppliedAmount:        "applied_amount2",
            TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            Memo:                 models.NewOptional(models.ToPointer("memo0")),
            Role:                 models.ToPointer("role0"),
            ConsolidatedInvoice:  models.ToPointer(false),
            AppliedCreditNotes:   []models.AppliedCreditNoteData{
                models.AppliedCreditNoteData{
                    Uid:                  models.ToPointer("uid4"),
                    Number:               models.ToPointer("number8"),
                },
                models.AppliedCreditNoteData{
                    Uid:                  models.ToPointer("uid4"),
                    Number:               models.ToPointer("number8"),
                },
                models.AppliedCreditNoteData{
                    Uid:                  models.ToPointer("uid4"),
                    Number:               models.ToPointer("number8"),
                },
            },
        },
    }

}
```

