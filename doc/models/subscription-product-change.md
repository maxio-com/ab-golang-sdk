
# Subscription Product Change

Event data for both `subscription_product_change` and `subscription_product_change_scheduled`. The price point and `effective_at` fields are only populated for scheduled changes.

## Structure

`SubscriptionProductChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousProductId` | `int` | Required | - |
| `NewProductId` | `int` | Required | - |
| `PreviousProductPricePointId` | `models.Optional[int]` | Optional | - |
| `NewProductPricePointId` | `models.Optional[int]` | Optional | - |
| `EffectiveAt` | `models.Optional[time.Time]` | Optional | When the scheduled product change takes effect (the subscription's next renewal). Only sent for `subscription_product_change_scheduled`. |

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
    subscriptionProductChange := models.SubscriptionProductChange{
        PreviousProductId:           104,
        NewProductId:                10,
        PreviousProductPricePointId: models.NewOptional(models.ToPointer(228)),
        NewProductPricePointId:      models.NewOptional(models.ToPointer(222)),
        EffectiveAt:                 models.NewOptional(models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }))),
    }

}
```

