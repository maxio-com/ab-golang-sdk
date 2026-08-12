
# Credit Note Application

## Structure

`CreditNoteApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `TransactionTime` | `*time.Time` | Optional | - |
| `InvoiceUid` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | - |
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
    creditNoteApplication := models.CreditNoteApplication{
        Uid:                  models.ToPointer("uid0"),
        TransactionTime:      models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        InvoiceUid:           models.ToPointer("invoice_uid0"),
        Memo:                 models.ToPointer("memo4"),
        AppliedAmount:        models.ToPointer("applied_amount8"),
    }

}
```

