
# Invoice Credit

## Structure

`InvoiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `CreditNoteNumber` | `*string` | Optional | - |
| `CreditNoteUid` | `*string` | Optional | - |
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
    invoiceCredit := models.InvoiceCredit{
        Uid:                  models.ToPointer("uid8"),
        CreditNoteNumber:     models.ToPointer("credit_note_number2"),
        CreditNoteUid:        models.ToPointer("credit_note_uid2"),
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Memo:                 models.ToPointer("memo2"),
    }

}
```

