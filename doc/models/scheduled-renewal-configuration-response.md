
# Scheduled Renewal Configuration Response

## Structure

`ScheduledRenewalConfigurationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ScheduledRenewalConfiguration` | [`*models.ScheduledRenewalConfiguration`](../../doc/models/scheduled-renewal-configuration.md) | Optional | - |

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
    scheduledRenewalConfigurationResponse := models.ScheduledRenewalConfigurationResponse{
        ScheduledRenewalConfiguration: models.ToPointer(models.ScheduledRenewalConfiguration{
            Id:                                 models.ToPointer(134),
            SiteId:                             models.ToPointer(60),
            SubscriptionId:                     models.ToPointer(244),
            StartsAt:                           models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            EndsAt:                             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        }),
    }

}
```

