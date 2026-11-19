
# Feature Templates List Response

## Structure

`FeatureTemplatesListResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Items` | [`[]models.FeatureTemplate`](../../doc/models/feature-template.md) | Required | - |
| `TotalCount` | `int` | Required | Total number of feature templates matching the filters, across all pages. |
| `ArchivedCount` | `int` | Required | Number of archived feature templates matching the filters. Returned as `0` unless the active result set is empty or `status=archived` was requested. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureTemplatesListResponse := models.FeatureTemplatesListResponse{
        Items:                []models.FeatureTemplate{
            models.FeatureTemplate{
                Id:                         models.ToPointer(66),
                Key:                        models.ToPointer("sso"),
                Name:                       models.ToPointer("Single Sign-On"),
                Description:                models.NewOptional(models.ToPointer("description2")),
                Kind:                       models.ToPointer(models.FeatureKind_ACCESSRIGHT),
            },
        },
        TotalCount:           54,
        ArchivedCount:        78,
    }

}
```

