
# Update Metafields Request

## Structure

`UpdateMetafieldsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metafields` | [`*models.UpdateMetafieldsRequestMetafields`](../../doc/models/containers/update-metafields-request-metafields.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateMetafieldsRequest := models.UpdateMetafieldsRequest{
        Metafields:           models.ToPointer(models.UpdateMetafieldsRequestMetafieldsContainer.FromUpdateMetafield(models.UpdateMetafield{
            CurrentName:          models.ToPointer("current_name0"),
            Name:                 models.ToPointer("name6"),
            Scope:                models.ToPointer(models.MetafieldScope{
                Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
                Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
                Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
                Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
                PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
            }),
            InputType:            models.ToPointer(models.MetafieldInput_BALANCETRACKER),
            Enum:                 []string{
                "enum2",
            },
        })),
    }

}
```

