
# Invoice Avatax Details

## Structure

`InvoiceAvataxDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[int64]` | Optional | - |
| `Status` | `models.Optional[string]` | Optional | - |
| `DocumentCode` | `models.Optional[string]` | Optional | - |
| `CommitDate` | `models.Optional[time.Time]` | Optional | - |
| `ModifyDate` | `models.Optional[time.Time]` | Optional | - |

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
    invoiceAvataxDetails := models.InvoiceAvataxDetails{
        Id:                   models.NewOptional(models.ToPointer(int64(184))),
        Status:               models.NewOptional(models.ToPointer("status2")),
        DocumentCode:         models.NewOptional(models.ToPointer("document_code4")),
        CommitDate:           models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
        ModifyDate:           models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
    }

}
```

