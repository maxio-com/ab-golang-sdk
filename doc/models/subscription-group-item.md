
# Subscription Group Item

## Structure

`SubscriptionGroupItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Reference` | `models.Optional[string]` | Optional | - |
| `CustomerId` | `*int` | Optional | The ID of the customer the subscription belongs to. This is the payer unless the signup request named another customer for the subscription. |
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
        CustomerId:              models.ToPointer(64),
        ProductId:               models.ToPointer(32),
        ProductHandle:           models.NewOptional(models.ToPointer("product_handle8")),
    }

}
```

