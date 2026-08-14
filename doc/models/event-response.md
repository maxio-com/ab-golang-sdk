
# Event Response

## Structure

`EventResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Event` | [`models.Event`](../../doc/models/event.md) | Required | - |

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
    eventResponse := models.EventResponse{
        Event:                models.Event{
            Id:                   int64(242),
            Key:                  models.EventKey_SUBSCRIPTIONREMOVEDFROMGROUP,
            Message:              "message0",
            SubscriptionId:       models.ToPointer(96),
            CustomerId:           models.ToPointer(24),
            CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
            EventSpecificData:    models.ToPointer(models.EventEventSpecificDataContainer.FromSubscriptionProductChange(models.SubscriptionProductChange{
                PreviousProductId:           126,
                NewProductId:                12,
                PreviousProductPricePointId: models.NewOptional(models.ToPointer(250)),
                NewProductPricePointId:      models.NewOptional(models.ToPointer(244)),
                EffectiveAt:                 models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
            })),
        },
    }

}
```

