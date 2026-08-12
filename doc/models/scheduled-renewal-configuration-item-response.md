
# Scheduled Renewal Configuration Item Response

## Structure

`ScheduledRenewalConfigurationItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ScheduledRenewalConfigurationItem` | [`*models.ScheduledRenewalConfigurationItem`](../../doc/models/scheduled-renewal-configuration-item.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalConfigurationItemResponse := models.ScheduledRenewalConfigurationItemResponse{
        ScheduledRenewalConfigurationItem: models.ToPointer(models.ScheduledRenewalConfigurationItem{
            Id:                                 models.ToPointer(98),
            SubscriptionId:                     models.ToPointer(208),
            SubscriptionRenewalConfigurationId: models.ToPointer(108),
            ItemId:                             models.ToPointer(246),
            ItemType:                           models.ToPointer("item_type2"),
        }),
    }

}
```

