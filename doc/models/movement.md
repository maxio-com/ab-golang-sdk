
# Movement

## Structure

`Movement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Timestamp` | `*time.Time` | Optional | - |
| `AmountInCents` | `*int64` | Optional | - |
| `AmountFormatted` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Category` | `*string` | Optional | - |
| `Breakouts` | [`*models.Breakouts`](../../doc/models/breakouts.md) | Optional | - |
| `LineItems` | [`[]models.MovementLineItem`](../../doc/models/movement-line-item.md) | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `SubscriberName` | `*string` | Optional | - |

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
    movement := models.Movement{
        Timestamp:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        AmountInCents:        models.ToPointer(int64(34)),
        AmountFormatted:      models.ToPointer("amount_formatted6"),
        Description:          models.ToPointer("description4"),
        Category:             models.ToPointer("category2"),
    }

}
```

