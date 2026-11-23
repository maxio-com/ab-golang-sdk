
# Clone Component Price Point Request

## Structure

`CloneComponentPricePointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | [`models.CloneComponentPricePoint`](../../doc/models/clone-component-price-point.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cloneComponentPricePointRequest := models.CloneComponentPricePointRequest{
        PricePoint:           models.CloneComponentPricePoint{
            Name:                 "name0",
            Handle:               models.ToPointer("handle6"),
            FeatureCatalogItems:  []models.CloneFeatureCatalogItem{
                models.CloneFeatureCatalogItem{
                    FeatureTemplateId:       8,
                    Value:                   "value4",
                    PeriodicityInterval:     models.NewOptional(models.ToPointer(136)),
                    PeriodicityUnit:         models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_HOUR)),
                    GrantExpirationInterval: models.NewOptional(models.ToPointer(60)),
                    GrantExpirationUnit:     models.NewOptional(models.ToPointer(models.GrantExpirationUnit_MONTH)),
                },
                models.CloneFeatureCatalogItem{
                    FeatureTemplateId:       8,
                    Value:                   "value4",
                    PeriodicityInterval:     models.NewOptional(models.ToPointer(136)),
                    PeriodicityUnit:         models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_HOUR)),
                    GrantExpirationInterval: models.NewOptional(models.ToPointer(60)),
                    GrantExpirationUnit:     models.NewOptional(models.ToPointer(models.GrantExpirationUnit_MONTH)),
                },
                models.CloneFeatureCatalogItem{
                    FeatureTemplateId:       8,
                    Value:                   "value4",
                    PeriodicityInterval:     models.NewOptional(models.ToPointer(136)),
                    PeriodicityUnit:         models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_HOUR)),
                    GrantExpirationInterval: models.NewOptional(models.ToPointer(60)),
                    GrantExpirationUnit:     models.NewOptional(models.ToPointer(models.GrantExpirationUnit_MONTH)),
                },
            },
        },
    }

}
```

