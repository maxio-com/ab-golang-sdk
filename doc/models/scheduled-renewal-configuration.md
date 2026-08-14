
# Scheduled Renewal Configuration

## Structure

`ScheduledRenewalConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | ID of the renewal. |
| `SiteId` | `*int` | Optional | ID of the site to which the renewal belongs. |
| `SubscriptionId` | `*int` | Optional | The id of the subscription. |
| `StartsAt` | `*time.Time` | Optional | - |
| `EndsAt` | `*time.Time` | Optional | - |
| `LockInAt` | `*time.Time` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `Status` | `*string` | Optional | - |
| `ScheduledRenewalConfigurationItems` | [`[]models.ScheduledRenewalConfigurationItem`](../../doc/models/scheduled-renewal-configuration-item.md) | Optional | - |
| `Contract` | [`*models.Contract`](../../doc/models/contract.md) | Optional | Contract linked to the scheduled renewal configuration. |

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
    scheduledRenewalConfiguration := models.ScheduledRenewalConfiguration{
        Id:                                 models.ToPointer(8),
        SiteId:                             models.ToPointer(190),
        SubscriptionId:                     models.ToPointer(118),
        StartsAt:                           models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndsAt:                             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

