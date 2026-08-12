
# Update Coupon Currency

## Structure

`UpdateCouponCurrency`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currency` | `string` | Required | ISO code for the site defined currency. |
| `Price` | `int` | Required | Price for the given currency. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateCouponCurrency := models.UpdateCouponCurrency{
        Currency:             "currency4",
        Price:                100,
    }

}
```

