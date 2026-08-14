
# Coupon Request

## Structure

`CouponRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coupon` | [`*models.CouponPayload`](../../doc/models/coupon-payload.md) | Optional | - |
| `RestrictedProducts` | `map[string]bool` | Optional | An object where the keys are product IDs or handles (prefixed with 'handle:'), and the values are booleans indicating if the coupon should be applicable to the product. |
| `RestrictedComponents` | `map[string]bool` | Optional | An object where the keys are component IDs or handles (prefixed with 'handle:'), and the values are booleans indicating if the coupon should be applicable to the component. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponRequest := models.CouponRequest{
        Coupon:               models.ToPointer(models.CouponPayload{
            Name:                          models.ToPointer("name4"),
            Code:                          models.ToPointer("code2"),
            Description:                   models.ToPointer("description6"),
            Percentage:                    models.ToPointer(models.CouponPayloadPercentageContainer.FromString("String3")),
            AmountInCents:                 models.ToPointer(int64(230)),
        }),
        RestrictedProducts:   map[string]bool{
            "key0": true,
            "key1": false,
        },
        RestrictedComponents: map[string]bool{
            "key0": true,
            "key1": false,
        },
    }

}
```

