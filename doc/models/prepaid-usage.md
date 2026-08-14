
# Prepaid Usage

## Structure

`PrepaidUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `PreviousOverageUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewUnitBalance` | [`models.PrepaidUsageNewUnitBalance`](../../doc/models/containers/prepaid-usage-new-unit-balance.md) | Required | This is a container for one-of cases. |
| `NewOverageUnitBalance` | [`models.PrepaidUsageNewOverageUnitBalance`](../../doc/models/containers/prepaid-usage-new-overage-unit-balance.md) | Required | This is a container for one-of cases. |
| `UsageQuantity` | `int` | Required | - |
| `OverageUsageQuantity` | `int` | Required | - |
| `ComponentId` | `int` | Required | - |
| `ComponentHandle` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `AllocationDetails` | [`[]models.PrepaidUsageAllocationDetail`](../../doc/models/prepaid-usage-allocation-detail.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidUsage := models.PrepaidUsage{
        PreviousUnitBalance:        "previous_unit_balance4",
        PreviousOverageUnitBalance: "previous_overage_unit_balance0",
        NewUnitBalance:             models.PrepaidUsageNewUnitBalanceContainer.FromNumber(206),
        NewOverageUnitBalance:      models.PrepaidUsageNewOverageUnitBalanceContainer.FromNumber(78),
        UsageQuantity:              246,
        OverageUsageQuantity:       138,
        ComponentId:                208,
        ComponentHandle:            "component_handle0",
        Memo:                       "memo4",
        AllocationDetails:          []models.PrepaidUsageAllocationDetail{
            models.PrepaidUsageAllocationDetail{
                AllocationId:         models.ToPointer(18),
                ChargeId:             models.ToPointer(84),
                UsageQuantity:        models.ToPointer(10),
            },
        },
    }

}
```

