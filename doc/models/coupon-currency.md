
# Coupon Currency

## Structure

`CouponCurrency`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[int]` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `Price` | `models.Optional[float64]` | Optional | - |
| `CouponId` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponCurrency := models.CouponCurrency{
        Id:                   models.NewOptional(models.ToPointer(202)),
        Currency:             models.ToPointer("currency0"),
        Price:                models.NewOptional(models.ToPointer(float64(14.62))),
        CouponId:             models.ToPointer(184),
    }

}
```

