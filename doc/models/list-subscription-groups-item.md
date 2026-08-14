
# List Subscription Groups Item

## Structure

`ListSubscriptionGroupsItem`

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
| `AccountBalances` | [`*models.SubscriptionGroupBalances`](../../doc/models/subscription-group-balances.md) | Optional | - |
| `GroupType` | [`*models.GroupType`](../../doc/models/group-type.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupsItem := models.ListSubscriptionGroupsItem{
        Uid:                   models.ToPointer("uid0"),
        Scheme:                models.ToPointer(228),
        CustomerId:            models.ToPointer(248),
        PaymentProfileId:      models.ToPointer(100),
        SubscriptionIds:       []int{
            102,
            103,
        },
    }

}
```

