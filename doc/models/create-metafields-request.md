
# Create Metafields Request

## Structure

`CreateMetafieldsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metafields` | [`models.CreateMetafieldsRequestMetafields`](../../doc/models/containers/create-metafields-request-metafields.md) | Required | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMetafieldsRequest := models.CreateMetafieldsRequest{
        Metafields:           models.CreateMetafieldsRequestMetafieldsContainer.FromCreateMetafield(models.CreateMetafield{
            Name:                 models.ToPointer("my_field"),
            Scope:                models.ToPointer(models.MetafieldScope{
                Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
                Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
                Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
                Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
                PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
                PublicEdit:           models.ToPointer(models.IncludeOption_EXCLUDE),
            }),
            InputType:            models.ToPointer(models.MetafieldInput_TEXT),
            Enum:                 []string{
                "string",
            },
        }),
    }

}
```

