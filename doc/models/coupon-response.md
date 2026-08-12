
# Coupon Response

## Structure

`CouponResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coupon` | [`*models.Coupon`](../../doc/models/coupon.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponResponse := models.CouponResponse{
        Coupon:               models.ToPointer(models.Coupon{
            Id:                            models.ToPointer(196),
            Name:                          models.ToPointer("name4"),
            Code:                          models.ToPointer("code2"),
            Description:                   models.ToPointer("description6"),
            Amount:                        models.NewOptional(models.ToPointer(float64(97.66))),
        }),
    }

}
```

