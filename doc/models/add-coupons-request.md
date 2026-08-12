
# Add Coupons Request

## Structure

`AddCouponsRequest`

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
    addCouponsRequest := models.AddCouponsRequest{
        Codes:                []string{
            "codes6",
            "codes7",
        },
    }

}
```

