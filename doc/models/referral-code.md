
# Referral Code

## Structure

`ReferralCode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Code` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    referralCode := models.ReferralCode{
        Id:                   models.ToPointer(186),
        SiteId:               models.ToPointer(112),
        SubscriptionId:       models.ToPointer(40),
        Code:                 models.ToPointer("code6"),
    }

}
```

