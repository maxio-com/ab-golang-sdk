
# Apply Debit Note Event Data

Example schema for an `apply_debit_note` event

## Structure

`ApplyDebitNoteEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DebitNoteNumber` | `string` | Required | A unique, identifying string that appears on the debit note and in places it is referenced. |
| `DebitNoteUid` | `string` | Required | Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters. |
| `OriginalAmount` | `string` | Required | The full, original amount of the debit note. |
| `AppliedAmount` | `string` | Required | The amount of the debit note applied to invoice. |
| `Memo` | `models.Optional[string]` | Optional | The debit note memo. |
| `TransactionTime` | `models.Optional[time.Time]` | Optional | The time the debit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |

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
    applyDebitNoteEventData := models.ApplyDebitNoteEventData{
        DebitNoteNumber:      "debit_note_number8",
        DebitNoteUid:         "debit_note_uid4",
        OriginalAmount:       "original_amount2",
        AppliedAmount:        "applied_amount0",
        Memo:                 models.NewOptional(models.ToPointer("memo2")),
        TransactionTime:      models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
    }

}
```

