
# Scheduled Renewal Lock in Request

## Structure

`ScheduledRenewalLockInRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LockInAt` | `time.Time` | Required | Date to lock in the renewal. |

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
    scheduledRenewalLockInRequest := models.ScheduledRenewalLockInRequest{
        LockInAt:             parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
    }

}
```

