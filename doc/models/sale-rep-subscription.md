
# Sale Rep Subscription

## Structure

`SaleRepSubscription`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SiteName` | `*string` | Optional | - |
| `SubscriptionUrl` | `*string` | Optional | - |
| `CustomerName` | `*string` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `Mrr` | `*string` | Optional | - |
| `Usage` | `*string` | Optional | - |
| `Recurring` | `*string` | Optional | - |
| `LastPayment` | `*string` | Optional | - |
| `ChurnDate` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    saleRepSubscription := models.SaleRepSubscription{
        Id:                   models.ToPointer(68),
        SiteName:             models.ToPointer("site_name8"),
        SubscriptionUrl:      models.ToPointer("subscription_url2"),
        CustomerName:         models.ToPointer("customer_name8"),
        CreatedAt:            models.ToPointer("created_at4"),
    }

}
```

