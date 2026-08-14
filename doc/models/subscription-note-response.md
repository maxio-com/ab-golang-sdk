
# Subscription Note Response

## Structure

`SubscriptionNoteResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | [`models.SubscriptionNote`](../../doc/models/subscription-note.md) | Required | - |

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
    subscriptionNoteResponse := models.SubscriptionNoteResponse{
        Note:                 models.SubscriptionNote{
            Id:                   models.ToPointer(28),
            Body:                 models.ToPointer("body0"),
            SubscriptionId:       models.ToPointer(138),
            CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            UpdatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        },
    }

}
```

