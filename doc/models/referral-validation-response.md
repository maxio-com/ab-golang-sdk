
# Referral Validation Response

## Structure

`ReferralValidationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ReferralCode` | [`*models.ReferralCode`](../../doc/models/referral-code.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    referralValidationResponse := models.ReferralValidationResponse{
        ReferralCode:         models.ToPointer(models.ReferralCode{
            Id:                   models.ToPointer(46),
            SiteId:               models.ToPointer(228),
            SubscriptionId:       models.ToPointer(156),
            Code:                 models.ToPointer("code0"),
        }),
    }

}
```

