
# Subscription Note

## Structure

`SubscriptionNote`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Body` | `*string` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `Sticky` | `*bool` | Optional | - |

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
    subscriptionNote := models.SubscriptionNote{
        Id:                   models.ToPointer(18),
        Body:                 models.ToPointer("body6"),
        SubscriptionId:       models.ToPointer(128),
        CreatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        UpdatedAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

