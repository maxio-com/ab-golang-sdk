
# Update Feature Template Request

## Structure

`UpdateFeatureTemplateRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.Feature1`](../../doc/models/feature-1.md) | Required | `key` cannot be changed once set. `kind` cannot be changed once any feature catalog item has been created from this template. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateFeatureTemplateRequest := models.UpdateFeatureTemplateRequest{
        Feature:              models.Feature1{
            Name:                       models.ToPointer("name4"),
            Description:                models.NewOptional(models.ToPointer("description4")),
            Unit:                       models.NewOptional(models.ToPointer("unit2")),
            ValueType:                  models.ToPointer(models.FeatureValueType_BOOLEAN),
            DefaultValue:               models.NewOptional(models.ToPointer("default_value4")),
        },
    }

}
```

