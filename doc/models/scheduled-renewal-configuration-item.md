
# Scheduled Renewal Configuration Item

## Structure

`ScheduledRenewalConfigurationItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `SubscriptionRenewalConfigurationId` | `*int` | Optional | - |
| `ItemId` | `*int` | Optional | - |
| `ItemType` | `*string` | Optional | - |
| `ItemSubclass` | `*string` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `PricePointType` | `*string` | Optional | - |
| `Quantity` | `*int` | Optional | - |
| `DecimalQuantity` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalConfigurationItem := models.ScheduledRenewalConfigurationItem{
        Id:                                 models.ToPointer(54),
        SubscriptionId:                     models.ToPointer(164),
        SubscriptionRenewalConfigurationId: models.ToPointer(64),
        ItemId:                             models.ToPointer(202),
        ItemType:                           models.ToPointer("item_type0"),
    }

}
```

