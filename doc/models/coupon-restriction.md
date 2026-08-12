
# Coupon Restriction

## Structure

`CouponRestriction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ItemType` | [`*models.RestrictionType`](../../doc/models/restriction-type.md) | Optional | - |
| `ItemId` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    couponRestriction := models.CouponRestriction{
        Id:                   models.ToPointer(66),
        ItemType:             models.ToPointer(models.RestrictionType_COMPONENT),
        ItemId:               models.ToPointer(214),
        Name:                 models.ToPointer("name0"),
        Handle:               models.NewOptional(models.ToPointer("handle6")),
    }

}
```

