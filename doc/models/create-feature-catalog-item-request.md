
# Create Feature Catalog Item Request

The owning product or component is taken from the URL and must not be included in the request body.

## Structure

`CreateFeatureCatalogItemRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.Feature2`](../../doc/models/feature-2.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createFeatureCatalogItemRequest := models.CreateFeatureCatalogItemRequest{
        Feature:              models.Feature2{
            FeatureTemplateId:        196,
            Value:                    models.ToPointer("value6"),
            PeriodicityInterval:      models.NewOptional(models.ToPointer(68)),
            PeriodicityUnit:          models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_WEEK)),
            PricePointType:           models.NewOptional(models.ToPointer(models.FeatureOwnerPricePointType_PRODUCTPRICEPOINT)),
            PricePointId:             models.NewOptional(models.ToPointer(84)),
            PropagateToSubscriptions: models.ToPointer(false),
        },
    }

}
```

