
# Coupon Currency Response

## Structure

`CouponCurrencyResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.CouponCurrency`](../../doc/models/coupon-currency.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponCurrencyResponse := models.CouponCurrencyResponse{
        CurrencyPrices:       []models.CouponCurrency{
            models.CouponCurrency{
                Id:                   models.NewOptional(models.ToPointer(50)),
                Currency:             models.ToPointer("currency8"),
                Price:                models.NewOptional(models.ToPointer(float64(233.74))),
                CouponId:             models.ToPointer(224),
            },
        },
    }

}
```

