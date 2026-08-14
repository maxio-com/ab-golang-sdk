
# Scheduled Renewal Configuration Request Body

## Structure

`ScheduledRenewalConfigurationRequestBody`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartsAt` | `*time.Time` | Optional | (Optional) Start of the renewal term. |
| `EndsAt` | `*time.Time` | Optional | (Optional) End of the renewal term. |
| `LockInAt` | `*time.Time` | Optional | (Optional) Lock-in date for the renewal. |
| `ContractId` | `*int` | Optional | (Optional) Existing contract to associate with the scheduled renewal. Contracts must be enabled for your site. |
| `CreateNewContract` | `*bool` | Optional | (Optional) Set to true to create a new contract when contracts are enabled. Contracts must be enabled for your site. |

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
    scheduledRenewalConfigurationRequestBody := models.ScheduledRenewalConfigurationRequestBody{
        StartsAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        EndsAt:               models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        LockInAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        ContractId:           models.ToPointer(162),
        CreateNewContract:    models.ToPointer(false),
    }

}
```

