
# Feature Template Response

## Structure

`FeatureTemplateResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.FeatureTemplate`](../../doc/models/feature-template.md) | Required | A feature that can be granted to subscribers, defined once at the site level and then attached to products or components. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureTemplateResponse := models.FeatureTemplateResponse{
        Feature:              models.FeatureTemplate{
            Id:                         models.ToPointer(206),
            Key:                        models.ToPointer("sso"),
            Name:                       models.ToPointer("Single Sign-On"),
            Description:                models.NewOptional(models.ToPointer("description4")),
            Kind:                       models.ToPointer(models.FeatureKind_USAGELIMIT),
        },
    }

}
```

