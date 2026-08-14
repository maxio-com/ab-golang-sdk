
# Offer Discount

## Structure

`OfferDiscount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CouponCode` | `*string` | Optional | - |
| `CouponId` | `*int` | Optional | - |
| `CouponName` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    offerDiscount := models.OfferDiscount{
        CouponCode:           models.ToPointer("coupon_code6"),
        CouponId:             models.ToPointer(202),
        CouponName:           models.ToPointer("coupon_name6"),
    }

}
```

