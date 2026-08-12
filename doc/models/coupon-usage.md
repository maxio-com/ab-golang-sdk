
# Coupon Usage

## Structure

`CouponUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify id of the product |
| `Name` | `*string` | Optional | Name of the product |
| `Signups` | `*int` | Optional | Number of times the coupon has been applied |
| `Savings` | `models.Optional[int]` | Optional | Dollar amount of customer savings as a result of the coupon. |
| `SavingsInCents` | `models.Optional[int64]` | Optional | Dollar amount of customer savings as a result of the coupon. |
| `Revenue` | `models.Optional[int]` | Optional | Total revenue of all subscriptions that have received a discount from this coupon. |
| `RevenueInCents` | `*int64` | Optional | Total revenue of all subscriptions that have received a discount from this coupon. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponUsage := models.CouponUsage{
        Id:                   models.ToPointer(240),
        Name:                 models.ToPointer("name8"),
        Signups:              models.ToPointer(4),
        Savings:              models.NewOptional(models.ToPointer(22)),
        SavingsInCents:       models.NewOptional(models.ToPointer(int64(108))),
    }

}
```

