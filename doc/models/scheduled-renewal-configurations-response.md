
# Scheduled Renewal Configurations Response

## Structure

`ScheduledRenewalConfigurationsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ScheduledRenewalConfigurations` | [`[]models.ScheduledRenewalConfiguration`](../../doc/models/scheduled-renewal-configuration.md) | Optional | - |

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
    scheduledRenewalConfigurationsResponse := models.ScheduledRenewalConfigurationsResponse{
        ScheduledRenewalConfigurations: []models.ScheduledRenewalConfiguration{
            models.ScheduledRenewalConfiguration{
                Id:                                 models.ToPointer(122),
                SiteId:                             models.ToPointer(48),
                SubscriptionId:                     models.ToPointer(232),
                StartsAt:                           models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
                EndsAt:                             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            },
        },
    }

}
```

