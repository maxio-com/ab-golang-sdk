
# Create Feature Template Request

## Structure

`CreateFeatureTemplateRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.Feature`](../../doc/models/feature.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createFeatureTemplateRequest := models.CreateFeatureTemplateRequest{
        Feature:              models.Feature{
            Key:                        "sso",
            Name:                       "name4",
            Description:                models.NewOptional(models.ToPointer("description4")),
            Kind:                       models.FeatureKind_USAGELIMIT,
            Unit:                       models.NewOptional(models.ToPointer("unit2")),
            ValueType:                  models.ToPointer(models.FeatureValueType_BOOLEAN),
            DefaultValue:               models.NewOptional(models.ToPointer("default_value4")),
            DefaultPeriodicityInterval: models.NewOptional(models.ToPointer(150)),
        },
    }

}
```

