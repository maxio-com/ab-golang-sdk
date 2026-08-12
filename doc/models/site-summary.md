
# Site Summary

## Structure

`SiteSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SellerName` | `*string` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `SiteCurrency` | `*string` | Optional | - |
| `Stats` | [`*models.SiteStatistics`](../../doc/models/site-statistics.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    siteSummary := models.SiteSummary{
        SellerName:           models.ToPointer("seller_name4"),
        SiteName:             models.ToPointer("site_name8"),
        SiteId:               models.ToPointer(252),
        SiteCurrency:         models.ToPointer("site_currency0"),
        Stats:                models.ToPointer(models.SiteStatistics{
            TotalSubscriptions:         models.ToPointer(110),
            SubscriptionsToday:         models.ToPointer(228),
            TotalRevenue:               models.ToPointer("total_revenue6"),
            RevenueToday:               models.ToPointer("revenue_today4"),
            RevenueThisMonth:           models.ToPointer("revenue_this_month4"),
        }),
    }

}
```

