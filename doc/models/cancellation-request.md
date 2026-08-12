
# Cancellation Request

## Structure

`CancellationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.CancellationOptions`](../../doc/models/cancellation-options.md) | Required | - |

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
    cancellationRequest := models.CancellationRequest{
        Subscription:         models.CancellationOptions{
            CancellationMessage:            models.ToPointer("cancellation_message2"),
            ReasonCode:                     models.ToPointer("reason_code8"),
            CancelAtEndOfPeriod:            models.ToPointer(false),
            ScheduledCancellationAt:        models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
            RefundPrepaymentAccountBalance: models.ToPointer(false),
        },
    }

}
```

