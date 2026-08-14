
# Coupon Subcodes Response

## Structure

`CouponSubcodesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedCodes` | `[]string` | Optional | - |
| `DuplicateCodes` | `[]string` | Optional | - |
| `InvalidCodes` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponSubcodesResponse := models.CouponSubcodesResponse{
        CreatedCodes:         []string{
            "created_codes7",
        },
        DuplicateCodes:       []string{
            "duplicate_codes8",
        },
        InvalidCodes:         []string{
            "invalid_codes4",
            "invalid_codes3",
        },
    }

}
```

