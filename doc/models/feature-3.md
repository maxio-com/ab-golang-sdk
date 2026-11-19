
# Feature 3

## Structure

`Feature3`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Value` | `*string` | Optional | - |
| `PeriodicityInterval` | `models.Optional[int]` | Optional | - |
| `PeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | - |
| `PropagateToSubscriptions` | `*bool` | Optional | When `true`, existing subscriptions on the item's price point are updated with the new `value`/periodicity by a background job shortly after. Has no effect on a product/component-level item. When `false` (default) and the item's price point already has subscribers, the update is made on a clone of the price point instead, and existing subscribers are unaffected.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    feature3 := models.Feature3{
        Value:                    models.ToPointer("value0"),
        PeriodicityInterval:      models.NewOptional(models.ToPointer(222)),
        PeriodicityUnit:          models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_DAY)),
        PropagateToSubscriptions: models.ToPointer(false),
    }

}
```

