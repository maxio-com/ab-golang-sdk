
# Clone Component Price Point

## Structure

`CloneComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `*string` | Optional | - |
| `FeatureCatalogItems` | [`[]models.CloneFeatureCatalogItem`](../../doc/models/clone-feature-catalog-item.md) | Optional | The features the clone gets. The clone does not copy the source price point's features, so send every feature it should have. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cloneComponentPricePoint := models.CloneComponentPricePoint{
        Name:                 "name4",
        Handle:               models.ToPointer("handle0"),
        FeatureCatalogItems:  []models.CloneFeatureCatalogItem{
            models.CloneFeatureCatalogItem{
                FeatureTemplateId:       8,
                Value:                   "value4",
                PeriodicityInterval:     models.NewOptional(models.ToPointer(136)),
                PeriodicityUnit:         models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_HOUR)),
                GrantExpirationInterval: models.NewOptional(models.ToPointer(60)),
                GrantExpirationUnit:     models.NewOptional(models.ToPointer(models.GrantExpirationUnit_MONTH)),
            },
        },
    }

}
```

