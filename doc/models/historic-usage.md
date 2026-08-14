
# Historic Usage

(Optional) For Event Based Components. If the `include=historic_usages` query param is provided, the last ten billing periods will be returned.

## Structure

`HistoricUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalUsageQuantity` | `*float64` | Optional | Total usage of a component for billing period |
| `BillingPeriodStartsAt` | `*time.Time` | Optional | Start date of billing period |
| `BillingPeriodEndsAt` | `*time.Time` | Optional | End date of billing period |

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
    historicUsage := models.HistoricUsage{
        TotalUsageQuantity:    models.ToPointer(float64(140.74)),
        BillingPeriodStartsAt: models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        BillingPeriodEndsAt:   models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

