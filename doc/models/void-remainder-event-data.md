
# Void Remainder Event Data

Example schema for an `void_remainder` event

## Structure

`VoidRemainderEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditNoteAttributes` | [`models.CreditNote`](../../doc/models/credit-note.md) | Required | - |
| `Memo` | `string` | Required | The memo provided during invoice remainder voiding. |
| `AppliedAmount` | `string` | Required | The amount of the void. |
| `TransactionTime` | `time.Time` | Required | The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z" |

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
    voidRemainderEventData := models.VoidRemainderEventData{
        CreditNoteAttributes: models.CreditNote{
            Uid:                  models.ToPointer("uid2"),
            SiteId:               models.ToPointer(72),
            CustomerId:           models.ToPointer(184),
            SubscriptionId:       models.ToPointer(0),
            Number:               models.ToPointer("number0"),
        },
        Memo:                 "memo6",
        AppliedAmount:        "applied_amount6",
        TransactionTime:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    }

}
```

