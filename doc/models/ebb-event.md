
# EBB Event

## Structure

`EBBEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chargify` | [`*models.ChargifyEBB`](../../doc/models/chargify-ebb.md) | Optional | - |

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
    ebbEvent := models.EBBEvent{
        Chargify:             models.ToPointer(models.ChargifyEBB{
            Timestamp:             models.ToPointer(parseTime(time.RFC3339, "2020-02-27T17:45:50-05:00", func(err error) { log.Fatalln(err) })),
            SubscriptionId:        models.ToPointer(1),
        }),
    }

}
```

