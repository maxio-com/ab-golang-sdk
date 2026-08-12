
# Coupon Subcodes

## Structure

`CouponSubcodes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Codes` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponSubcodes := models.CouponSubcodes{
        Codes:                []string{
            "codes8",
            "codes9",
        },
    }

}
```

