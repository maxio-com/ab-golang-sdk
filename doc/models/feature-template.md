
# Feature Template

A feature that can be granted to subscribers, defined once at the site level and then attached to products or components.

## Structure

`FeatureTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Advanced Billing id of the feature template. |
| `Key` | `*string` | Optional | A unique, lowercase, underscore-separated identifier for the feature. Immutable once set. |
| `Name` | `*string` | Optional | The display name of the feature. |
| `Description` | `models.Optional[string]` | Optional | - |
| `Kind` | [`*models.FeatureKind`](../../doc/models/feature-kind.md) | Optional | The behavior of a feature:<br><br>- `access_right`: a boolean entitlement. A subscriber either has access or does not.<br>- `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").<br>- `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit. |
| `Unit` | `models.Optional[string]` | Optional | The unit the feature is measured in (for example, `requests` or `GB`). Required when `kind` is `usage_limit`. |
| `ValueType` | [`*models.FeatureValueType`](../../doc/models/feature-value-type.md) | Optional | The data type of a feature's value. For `access_right` features this is always `boolean`, and for `usage_limit` features this is always `numeric`. For `service_right` features, you choose the value type explicitly. |
| `DefaultValue` | `models.Optional[string]` | Optional | A default value used to pre-populate new feature catalog items created from this template. |
| `DefaultPeriodicityInterval` | `models.Optional[int]` | Optional | For `usage_limit` features, the default periodicity interval used to pre-populate new feature catalog items. Always `null` for other kinds. |
| `DefaultPeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | For `usage_limit` features, the default periodicity unit used to pre-populate new feature catalog items. Always `null` for other kinds. |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | The date and time the feature template was archived, or `null` if it is active. |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `ProductsCount` | `*int` | Optional | The number of **components** this feature template is currently attached to via an active feature catalog item. Despite the name, this counts components, not products. In the Advanced Billing UI, components are labeled "Products." |
| `PlansCount` | `*int` | Optional | The number of **products** this feature template is currently attached to via an active feature catalog item. Despite the name, this counts products, not plans. In the Advanced Billing UI, products are labeled "Plans." |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureTemplate := models.FeatureTemplate{
        Id:                         models.ToPointer(164),
        Key:                        models.ToPointer("sso"),
        Name:                       models.ToPointer("Single Sign-On"),
        Description:                models.NewOptional(models.ToPointer("description6")),
        Kind:                       models.ToPointer(models.FeatureKind_SERVICERIGHT),
    }

}
```

