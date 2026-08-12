
# Dunner Data

## Structure

`DunnerData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `State` | `string` | Required | - |
| `SubscriptionId` | `int` | Required | - |
| `RevenueAtRiskInCents` | `int64` | Required | - |
| `CreatedAt` | `time.Time` | Required | - |
| `Attempts` | `int` | Required | - |
| `LastAttemptedAt` | `time.Time` | Required | - |

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
    dunnerData := models.DunnerData{
        State:                "state4",
        SubscriptionId:       230,
        RevenueAtRiskInCents: int64(134),
        CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        Attempts:             6,
        LastAttemptedAt:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    }

}
```

