
# Billing Schedule

Billing schedule settings for component allocations or usages on multi-frequency subscriptions. Use this to start a component's billing period on a custom date instead of aligning with the product charge schedule.

## Structure

`BillingSchedule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InitialBillingAt` | `models.Optional[time.Time]` | Optional | Custom start date (ISO 8601 date, YYYY-MM-DD) for the component's first billing period. If omitted or null, billing aligns with the product schedule. If provided, date must be on or after the minimum allowed date for the subscription or component. |

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
    billingSchedule := models.BillingSchedule{
        InitialBillingAt:     models.NewOptional(models.ToPointer(parseTime(models.DEFAULT_DATE, "2026-01-01", func(err error) { log.Fatalln(err) }))),
    }

}
```

