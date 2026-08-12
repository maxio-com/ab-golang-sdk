
# Reactivate Subscription Group Response

## Structure

`ReactivateSubscriptionGroupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Scheme` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `PaymentProfileId` | `*int` | Optional | - |
| `SubscriptionIds` | `[]int` | Optional | - |
| `PrimarySubscriptionId` | `*int` | Optional | - |
| `NextAssessmentAt` | `*time.Time` | Optional | - |
| `State` | `*string` | Optional | - |
| `CancelAtEndOfPeriod` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reactivateSubscriptionGroupResponse := models.ReactivateSubscriptionGroupResponse{
        Uid:                   models.ToPointer("uid0"),
        Scheme:                models.ToPointer(10),
        CustomerId:            models.ToPointer(30),
        PaymentProfileId:      models.ToPointer(62),
        SubscriptionIds:       []int{
            140,
            141,
            142,
        },
    }

}
```

