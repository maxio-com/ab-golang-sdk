
# Webhook Response

## Structure

`WebhookResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Webhook` | [`*models.Webhook`](../../doc/models/webhook.md) | Optional | - |

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
    webhookResponse := models.WebhookResponse{
        Webhook:              models.ToPointer(models.Webhook{
            Event:                models.ToPointer("event2"),
            Id:                   models.ToPointer(int64(18)),
            CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            LastError:            models.ToPointer("last_error4"),
            LastErrorAt:          models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        }),
    }

}
```

