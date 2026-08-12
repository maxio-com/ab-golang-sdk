
# Webhook

## Structure

`Webhook`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Event` | `*string` | Optional | A string describing which event type produced the given webhook |
| `Id` | `*int64` | Optional | The unique identifier for the webhook (unique across all of Chargify). This is not changed on a retry/replay of the same webhook, so it may be used to avoid duplicate action for the same event. |
| `CreatedAt` | `*time.Time` | Optional | Timestamp indicating when the webhook was created |
| `LastError` | `*string` | Optional | Text describing the status code and/or error from the last failed attempt to send the Webhook. When a webhook is retried and accepted, this field will be cleared. |
| `LastErrorAt` | `*time.Time` | Optional | Timestamp indicating when the last non-acceptance occurred. If a webhook is later resent and accepted, this field will be cleared. |
| `AcceptedAt` | `models.Optional[time.Time]` | Optional | Timestamp indicating when the webhook was accepted by the merchant endpoint. When a webhook is explicitly replayed by the merchant, this value will be cleared until it is accepted again. |
| `LastSentAt` | `*time.Time` | Optional | Timestamp indicating when the most recent attempt was made to send the webhook |
| `LastSentUrl` | `*string` | Optional | The url that the endpoint was last sent to. |
| `Successful` | `*bool` | Optional | “A boolean flag describing whether the webhook was accepted by the webhook endpoint for the most recent attempt. (Acceptance is defined by receiving a “200 OK” HTTP response within a reasonable timeframe, e.g., 15 seconds.)” |
| `Body` | `*string` | Optional | The data sent within the webhook post |
| `Signature` | `*string` | Optional | The calculated webhook signature |
| `SignatureHmacSha256` | `*string` | Optional | The calculated HMAC-SHA-256 webhook signature |

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
    webhook := models.Webhook{
        Event:                models.ToPointer("event2"),
        Id:                   models.ToPointer(int64(18)),
        CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        LastError:            models.ToPointer("last_error4"),
        LastErrorAt:          models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

