
# Feature Catalog Items List Response

## Structure

`FeatureCatalogItemsListResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Features` | [`[]models.FeatureCatalogItem`](../../doc/models/feature-catalog-item.md) | Required | - |
| `SubscriptionsCount` | `int` | Required | The number of subscriptions on this product/component that would be affected if a feature catalog item change were propagated with `propagate_to_subscriptions=true`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureCatalogItemsListResponse := models.FeatureCatalogItemsListResponse{
        Features:             []models.FeatureCatalogItem{
            models.FeatureCatalogItem{
                Id:                   models.ToPointer(174),
                FeatureTemplateId:    models.ToPointer(164),
                FeatureKey:           models.ToPointer("feature_key2"),
                FeatureName:          models.ToPointer("feature_name6"),
                FeatureKind:          models.ToPointer(models.FeatureKind_ACCESSRIGHT),
            },
        },
        SubscriptionsCount:   22,
    }

}
```

