
# Feature Catalog Item

A feature template attached to a specific product or component (or one of their price points), with a concrete value. When a subscriber signs up for or is assigned this product/component, the feature catalog item is provisioned as an entitlement on their subscription.

## Structure

`FeatureCatalogItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FeatureTemplateId` | `*int` | Optional | The id of the feature template this item was created from. |
| `FeatureKey` | `*string` | Optional | The `key` of the parent feature template. |
| `FeatureName` | `*string` | Optional | The `name` of the parent feature template. |
| `FeatureKind` | [`*models.FeatureKind`](../../doc/models/feature-kind.md) | Optional | The behavior of a feature:<br><br>- `access_right`: a boolean entitlement. A subscriber either has access or does not.<br>- `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").<br>- `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit. |
| `Value` | `*string` | Optional | The value granted by this feature catalog item. Interpreted according to `feature_kind`: `"true"`/`"false"` for `access_right`, a numeric string for `usage_limit`, or any string for `service_right` (shaped by the feature template's `value_type`). |
| `PeriodicityInterval` | `models.Optional[int]` | Optional | Set when `feature_kind` is `usage_limit`; `null` otherwise. |
| `PeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | Set when `feature_kind` is `usage_limit`; `null` otherwise. |
| `PricePointType` | [`models.Optional[models.FeatureOwnerPricePointType]`](../../doc/models/feature-owner-price-point-type.md) | Optional | `null` when this feature catalog item applies to every price point of its owning product/component. Set when the feature catalog item is an override for one specific price point. |
| `PricePointId` | `models.Optional[int]` | Optional | Set together with `price_point_type` for price-point-specific overrides. |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureCatalogItem := models.FeatureCatalogItem{
        Id:                   models.ToPointer(52),
        FeatureTemplateId:    models.ToPointer(42),
        FeatureKey:           models.ToPointer("feature_key6"),
        FeatureName:          models.ToPointer("feature_name0"),
        FeatureKind:          models.ToPointer(models.FeatureKind_USAGELIMIT),
    }

}
```

