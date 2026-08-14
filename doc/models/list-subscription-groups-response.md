
# List Subscription Groups Response

## Structure

`ListSubscriptionGroupsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroups` | [`[]models.ListSubscriptionGroupsItem`](../../doc/models/list-subscription-groups-item.md) | Optional | - |
| `Meta` | [`*models.ListSubscriptionGroupsMeta`](../../doc/models/list-subscription-groups-meta.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionGroupsResponse := models.ListSubscriptionGroupsResponse{
        SubscriptionGroups:   []models.ListSubscriptionGroupsItem{
            models.ListSubscriptionGroupsItem{
                Uid:                   models.ToPointer("uid2"),
                Scheme:                models.ToPointer(166),
                CustomerId:            models.ToPointer(186),
                PaymentProfileId:      models.ToPointer(162),
                SubscriptionIds:       []int{
                    40,
                },
            },
        },
        Meta:                 models.ToPointer(models.ListSubscriptionGroupsMeta{
            CurrentPage:          models.ToPointer(126),
            TotalCount:           models.ToPointer(150),
        }),
    }

}
```

