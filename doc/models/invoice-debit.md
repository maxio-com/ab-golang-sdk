
# Invoice Debit

## Structure

`InvoiceDebit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `DebitNoteNumber` | `*string` | Optional | - |
| `DebitNoteUid` | `*string` | Optional | - |
| `Role` | [`*models.DebitNoteRole`](../../doc/models/debit-note-role.md) | Optional | The role of the debit note. |
| `TransactionTime` | `*time.Time` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |

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
    invoiceDebit := models.InvoiceDebit{
        Uid:                  models.ToPointer("uid8"),
        DebitNoteNumber:      models.ToPointer("debit_note_number8"),
        DebitNoteUid:         models.ToPointer("debit_note_uid4"),
        Role:                 models.ToPointer(models.DebitNoteRole_CHARGEBACK),
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

