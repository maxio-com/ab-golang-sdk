
# Sale Rep Settings

## Structure

`SaleRepSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomerName` | `*string` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `SiteLink` | `*string` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `SubscriptionMrr` | `*string` | Optional | - |
| `SalesRepId` | `*int` | Optional | - |
| `SalesRepName` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    saleRepSettings := models.SaleRepSettings{
        CustomerName:         models.ToPointer("customer_name6"),
        SubscriptionId:       models.ToPointer(202),
        SiteLink:             models.ToPointer("site_link2"),
        SiteName:             models.ToPointer("site_name6"),
        SubscriptionMrr:      models.ToPointer("subscription_mrr4"),
    }

}
```

