
# Feature

## Structure

`Feature`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Key` | `string` | Required | A unique, lowercase, underscore-separated identifier for the feature. Immutable once set. |
| `Name` | `string` | Required | The display name of the feature. |
| `Description` | `models.Optional[string]` | Optional | - |
| `Kind` | [`models.FeatureKind`](../../doc/models/feature-kind.md) | Required | The behavior of a feature:<br><br>- `access_right`: a boolean entitlement. A subscriber either has access or does not.<br>- `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").<br>- `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit. |
| `Unit` | `models.Optional[string]` | Optional | Required when `kind` is `usage_limit`. |
| `ValueType` | [`*models.FeatureValueType`](../../doc/models/feature-value-type.md) | Optional | Required when `kind` is `service_right`. Ignored for other kinds, where it is inferred automatically. |
| `DefaultValue` | `models.Optional[string]` | Optional | - |
| `DefaultPeriodicityInterval` | `models.Optional[int]` | Optional | Only valid when `kind` is `usage_limit`. |
| `DefaultPeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | Only valid when `kind` is `usage_limit`. Must be set together with `default_periodicity_interval`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    feature := models.Feature{
        Key:                        "sso",
        Name:                       "name4",
        Description:                models.NewOptional(models.ToPointer("description4")),
        Kind:                       models.FeatureKind_USAGELIMIT,
        Unit:                       models.NewOptional(models.ToPointer("unit2")),
        ValueType:                  models.ToPointer(models.FeatureValueType_BOOLEAN),
        DefaultValue:               models.NewOptional(models.ToPointer("default_value4")),
        DefaultPeriodicityInterval: models.NewOptional(models.ToPointer(150)),
    }

}
```

