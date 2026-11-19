
# Update Feature Catalog Item Request

## Structure

`UpdateFeatureCatalogItemRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Feature` | [`models.Feature3`](../../doc/models/feature-3.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateFeatureCatalogItemRequest := models.UpdateFeatureCatalogItemRequest{
        Feature:              models.Feature3{
            Value:                    models.ToPointer("value6"),
            PeriodicityInterval:      models.NewOptional(models.ToPointer(68)),
            PeriodicityUnit:          models.NewOptional(models.ToPointer(models.EntitlementPeriodicityUnit_WEEK)),
            PropagateToSubscriptions: models.ToPointer(false),
        },
    }

}
```

