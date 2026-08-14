
# Apply Credit Note Event Data

Example schema for an `apply_credit_note` event

## Structure

`ApplyCreditNoteEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | Unique identifier for the credit note application. It is generated automatically by Chargify and has the prefix "cdt_" followed by alphanumeric characters. |
| `CreditNoteNumber` | `string` | Required | A unique, identifying string that appears on the credit note and in places it is referenced. |
| `CreditNoteUid` | `string` | Required | Unique identifier for the credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters. |
| `OriginalAmount` | `string` | Required | The full, original amount of the credit note. |
| `AppliedAmount` | `string` | Required | The amount of the credit note applied to invoice. |
| `TransactionTime` | `*time.Time` | Optional | The time the credit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |
| `Memo` | `models.Optional[string]` | Optional | The credit note memo. |
| `Role` | `*string` | Optional | The role of the credit note (e.g. 'general') |
| `ConsolidatedInvoice` | `*bool` | Optional | Shows whether it was applied to consolidated invoice or not. |
| `AppliedCreditNotes` | [`[]models.AppliedCreditNoteData`](../../doc/models/applied-credit-note-data.md) | Optional | List of credit notes applied to children invoices (if consolidated invoice) |

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
    applyCreditNoteEventData := models.ApplyCreditNoteEventData{
        Uid:                  "uid0",
        CreditNoteNumber:     "credit_note_number6",
        CreditNoteUid:        "credit_note_uid4",
        OriginalAmount:       "original_amount4",
        AppliedAmount:        "applied_amount8",
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Memo:                 models.NewOptional(models.ToPointer("memo4")),
        Role:                 models.ToPointer("role4"),
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
    }

}
```

