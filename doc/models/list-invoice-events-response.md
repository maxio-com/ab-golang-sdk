
# List Invoice Events Response

## Structure

`ListInvoiceEventsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.InvoiceEvent`](../../doc/models/containers/invoice-event.md) | Optional | - |
| `Page` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |

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
    listInvoiceEventsResponse := models.ListInvoiceEventsResponse{
        Events:               []models.InvoiceEvent{
            models.InvoiceEventContainer.FromApplyCreditNoteEvent(models.ApplyCreditNoteEvent{
                Id:                   int64(214),
                Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
                Invoice:              models.Invoice{
                    Id:                         models.ToPointer(int64(166)),
                    Uid:                        models.ToPointer("uid6"),
                    SiteId:                     models.ToPointer(92),
                    CustomerId:                 models.ToPointer(204),
                    SubscriptionId:             models.ToPointer(20),
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
            }),
            models.InvoiceEventContainer.FromApplyCreditNoteEvent(models.ApplyCreditNoteEvent{
                Id:                   int64(214),
                Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
                Invoice:              models.Invoice{
                    Id:                         models.ToPointer(int64(166)),
                    Uid:                        models.ToPointer("uid6"),
                    SiteId:                     models.ToPointer(92),
                    CustomerId:                 models.ToPointer(204),
                    SubscriptionId:             models.ToPointer(20),
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
            }),
            models.InvoiceEventContainer.FromApplyCreditNoteEvent(models.ApplyCreditNoteEvent{
                Id:                   int64(214),
                Timestamp:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
                Invoice:              models.Invoice{
                    Id:                         models.ToPointer(int64(166)),
                    Uid:                        models.ToPointer("uid6"),
                    SiteId:                     models.ToPointer(92),
                    CustomerId:                 models.ToPointer(204),
                    SubscriptionId:             models.ToPointer(20),
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
            }),
        },
        Page:                 models.ToPointer(28),
        PerPage:              models.ToPointer(196),
        TotalPages:           models.ToPointer(94),
    }

}
```

