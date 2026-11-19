
# Feature Catalog Item Response

## Structure

`FeatureCatalogItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.FeatureCatalogItem`](../../doc/models/feature-catalog-item.md) | Required | A feature template attached to a specific product or component (or one of their price points), with a concrete value. When a subscriber signs up for or is assigned this product/component, the feature catalog item is provisioned as an entitlement on their subscription. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureCatalogItemResponse := models.FeatureCatalogItemResponse{
        Feature:              models.FeatureCatalogItem{
            Id:                   models.ToPointer(206),
            FeatureTemplateId:    models.ToPointer(196),
            FeatureKey:           models.ToPointer("feature_key2"),
            FeatureName:          models.ToPointer("feature_name6"),
            FeatureKind:          models.ToPointer(models.FeatureKind_SERVICERIGHT),
        },
    }

}
```

