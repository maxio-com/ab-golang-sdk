
# Coupon Currency Request

## Structure

`CouponCurrencyRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.UpdateCouponCurrency`](../../doc/models/update-coupon-currency.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponCurrencyRequest := models.CouponCurrencyRequest{
        CurrencyPrices:       []models.UpdateCouponCurrency{
            models.UpdateCouponCurrency{
                Currency:             "currency8",
                Price:                78,
            },
        },
    }

}
```

