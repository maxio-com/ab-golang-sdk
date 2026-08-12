
# Event

## Structure

`Event`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int64` | Required | - |
| `Key` | [`models.EventKey`](../../doc/models/event-key.md) | Required | - |
| `Message` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `SubscriptionId` | `*int` | Required | - |
| `CustomerId` | `*int` | Required | - |
| `CreatedAt` | `time.Time` | Required | - |
| `EventSpecificData` | [`*models.EventEventSpecificData`](../../doc/models/containers/event-event-specific-data.md) | Required | This is a container for one-of cases. |

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
    event := models.Event{
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
    }

}
```

