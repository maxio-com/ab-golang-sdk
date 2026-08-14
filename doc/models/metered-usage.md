
# Metered Usage

## Structure

`MeteredUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousUnitBalance` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewUnitBalance` | [`models.MeteredUsageNewUnitBalance`](../../doc/models/containers/metered-usage-new-unit-balance.md) | Required | This is a container for one-of cases. |
| `UsageQuantity` | `int` | Required | - |
| `ComponentId` | `int` | Required | - |
| `ComponentHandle` | `string` | Required | - |
| `Memo` | `string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    meteredUsage := models.MeteredUsage{
        PreviousUnitBalance:  "previous_unit_balance6",
        NewUnitBalance:       models.MeteredUsageNewUnitBalanceContainer.FromNumber(66),
        UsageQuantity:        106,
        ComponentId:          68,
        ComponentHandle:      "component_handle0",
        Memo:                 "memo4",
    }

}
```

