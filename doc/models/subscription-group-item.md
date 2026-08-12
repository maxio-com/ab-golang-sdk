
# Subscription Group Item

## Structure

`SubscriptionGroupItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Reference` | `models.Optional[string]` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductHandle` | `models.Optional[string]` | Optional | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `ProductPricePointHandle` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `CouponCode` | `models.Optional[string]` | Optional | - |
| `TotalRevenueInCents` | `*int64` | Optional | - |
| `BalanceInCents` | `*int64` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupItem := models.SubscriptionGroupItem{
        Id:                      models.ToPointer(26),
        Reference:               models.NewOptional(models.ToPointer("reference4")),
        ProductId:               models.ToPointer(32),
        ProductHandle:           models.NewOptional(models.ToPointer("product_handle8")),
        ProductPricePointId:     models.ToPointer(148),
    }

}
```

