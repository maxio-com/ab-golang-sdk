
# Aggregated Entitlement Periodicity

## Structure

`AggregatedEntitlementPeriodicity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `*int` | Optional | - |
| `Unit` | [`*models.EntitlementPeriodicityUnit`](../../doc/models/entitlement-periodicity-unit.md) | Optional | The recurring window over which a `usage_limit` feature's allowance resets. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    aggregatedEntitlementPeriodicity := models.AggregatedEntitlementPeriodicity{
        Interval:             models.ToPointer(128),
        Unit:                 models.ToPointer(models.EntitlementPeriodicityUnit_YEAR),
    }

}
```

