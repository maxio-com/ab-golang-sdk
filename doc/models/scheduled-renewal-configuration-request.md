
# Scheduled Renewal Configuration Request

## Structure

`ScheduledRenewalConfigurationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalConfiguration` | [`models.ScheduledRenewalConfigurationRequestBody`](../../doc/models/scheduled-renewal-configuration-request-body.md) | Required | - |

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
    scheduledRenewalConfigurationRequest := models.ScheduledRenewalConfigurationRequest{
        RenewalConfiguration: models.ScheduledRenewalConfigurationRequestBody{
            StartsAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            EndsAt:               models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            LockInAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            ContractId:           models.ToPointer(244),
            CreateNewContract:    models.ToPointer(false),
        },
    }

}
```

