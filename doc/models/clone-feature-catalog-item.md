
# Clone Feature Catalog Item

A feature to put on the cloned price point. Requires Entitlements to be enabled for the site; ignored otherwise.

Unlike Create Feature Catalog Item, nothing falls back to the feature template's defaults: send `value`, and for a `usage_limit` feature its periodicity (and any grant expiration) explicitly. If any item can't be applied, for example a missing `value`, a `usage_limit` item without periodicity, or an archived or unknown feature template, the whole clone fails with `422` and nothing is created.

## Structure

`CloneFeatureCatalogItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FeatureTemplateId` | `int` | Required | - |
| `Value` | `string` | Required | - |
| `PeriodicityInterval` | `models.Optional[int]` | Optional | Required for `usage_limit` features; must be omitted for other kinds. |
| `PeriodicityUnit` | [`models.Optional[models.EntitlementPeriodicityUnit]`](../../doc/models/entitlement-periodicity-unit.md) | Optional | Required for `usage_limit` features; must be omitted for other kinds. |
| `GrantExpirationInterval` | `models.Optional[int]` | Optional | `usage_limit` features only. A positive number of `grant_expiration_unit`s; required with `day` or `month`, and must be omitted with `until_next_refill`. |
| `GrantExpirationUnit` | [`models.Optional[models.GrantExpirationUnit]`](../../doc/models/grant-expiration-unit.md) | Optional | `usage_limit` features only. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cloneFeatureCatalogItem := models.CloneFeatureCatalogItem{
        FeatureTemplateId:       212,
        Value:                   "value6",
        PeriodicityInterval:     models.NewOptional(models.ToPointer(84)),
        PeriodicityUnit:         models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_WEEK)),
        GrantExpirationInterval: models.NewOptional(models.ToPointer(8)),
        GrantExpirationUnit:     models.NewOptional(models.ToPointer(models.GrantExpirationUnit_UNTILNEXTREFILL)),
    }

}
```

