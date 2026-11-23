
# Feature 2

## Structure

`Feature2`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FeatureTemplateId` | `int` | Required | The id of the feature template to attach. |
| `Value` | `*string` | Optional | The value granted. Defaults to the feature template's `default_value`; required when the template has none. |
| `PeriodicityInterval` | `models.Optional[int]` | Optional | Required for `usage_limit` features, unless the feature template has a default periodicity, which is used when omitted. Must be omitted for other kinds. |
| `PeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | Required for `usage_limit` features, unless the feature template has a default periodicity, which is used when omitted. Must be omitted for other kinds. |
| `PricePointType` | [`models.Optional[models.FeatureOwnerPricePointType]`](../../doc/models/feature-owner-price-point-type.md) | Optional | Set together with `price_point_id` to scope the feature catalog item to a single price point. Omit both to create a product/component-level item, which is a template: it is copied onto price points created afterwards (except bulk-created ones; Clone Component Price Point applies only the items sent with it), but not added to price points that already exist or to their subscribers. |
| `PricePointId` | `models.Optional[int]` | Optional | - |
| `PropagateToSubscriptions` | `*bool` | Optional | Applies to price-point-scoped items only. When `true`, existing subscriptions on that price point are granted an entitlement for this feature by a background job shortly after, instead of at their next subscription change. When `false` (default) and the price point already has subscribers, the feature is created on a clone of the price point instead, and existing subscribers are unaffected. On a product/component-level item it has no effect, except when the call restores that item's archived price-point copies (creating an item again for a feature template whose item was archived): then `true` grants the feature to those price points' subscribers, and without it the call returns `422` if any of them lack the feature.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    feature2 := models.Feature2{
        FeatureTemplateId:        90,
        Value:                    models.ToPointer("value4"),
        PeriodicityInterval:      models.NewOptional(models.ToPointer(218)),
        PeriodicityUnit:          models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_HOUR)),
        PricePointType:           models.NewOptional(models.ToPointer(models.FeatureOwnerPricePointType_PRODUCTPRICEPOINT)),
        PricePointId:             models.NewOptional(models.ToPointer(22)),
        PropagateToSubscriptions: models.ToPointer(false),
    }

}
```

