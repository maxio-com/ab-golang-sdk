
# Feature 1

`key` cannot be changed once set. `kind` cannot be changed once any feature catalog item has been created from this template.

## Structure

`Feature1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Unit` | `models.Optional[string]` | Optional | - |
| `ValueType` | [`*models.FeatureValueType`](../../doc/models/feature-value-type.md) | Optional | The data type of a feature's value. For `access_right` features this is always `boolean`, and for `usage_limit` features this is always `numeric`. For `service_right` features, you choose the value type explicitly. |
| `DefaultValue` | `models.Optional[string]` | Optional | - |
| `DefaultPeriodicityInterval` | `models.Optional[int]` | Optional | - |
| `DefaultPeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    feature1 := models.Feature1{
        Name:                       models.ToPointer("name8"),
        Description:                models.NewOptional(models.ToPointer("description8")),
        Unit:                       models.NewOptional(models.ToPointer("unit6")),
        ValueType:                  models.ToPointer(models.FeatureValueType_TEXT),
        DefaultValue:               models.NewOptional(models.ToPointer("default_value8")),
    }

}
```

