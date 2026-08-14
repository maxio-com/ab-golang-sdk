
# Site Statistics

## Structure

`SiteStatistics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalSubscriptions` | `*int` | Optional | - |
| `SubscriptionsToday` | `*int` | Optional | - |
| `TotalRevenue` | `*string` | Optional | - |
| `RevenueToday` | `*string` | Optional | - |
| `RevenueThisMonth` | `*string` | Optional | - |
| `RevenueThisYear` | `*string` | Optional | - |
| `TotalCanceledSubscriptions` | `*int` | Optional | - |
| `TotalActiveSubscriptions` | `*int` | Optional | - |
| `TotalPastDueSubscriptions` | `*int` | Optional | - |
| `TotalUnpaidSubscriptions` | `*int` | Optional | - |
| `TotalDunningSubscriptions` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    siteStatistics := models.SiteStatistics{
        TotalSubscriptions:         models.ToPointer(168),
        SubscriptionsToday:         models.ToPointer(170),
        TotalRevenue:               models.ToPointer("total_revenue2"),
        RevenueToday:               models.ToPointer("revenue_today0"),
        RevenueThisMonth:           models.ToPointer("revenue_this_month0"),
    }

}
```

