
# List Metafields Response

## Structure

`ListMetafieldsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalCount` | `*int` | Optional | - |
| `CurrentPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `Metafields` | [`[]models.Metafield`](../../doc/models/metafield.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listMetafieldsResponse := models.ListMetafieldsResponse{
        TotalCount:           models.ToPointer(228),
        CurrentPage:          models.ToPointer(204),
        TotalPages:           models.ToPointer(216),
        PerPage:              models.ToPointer(74),
        Metafields:           []models.Metafield{
            models.Metafield{
                Id:                   models.ToPointer(22),
                Name:                 models.ToPointer("name2"),
                Scope:                models.ToPointer(models.MetafieldScope{
                    Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
                    Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
                    Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
                    Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
                    PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
                }),
                DataCount:            models.ToPointer(10),
                InputType:            models.ToPointer(models.MetafieldInput_BALANCETRACKER),
            },
            models.Metafield{
                Id:                   models.ToPointer(22),
                Name:                 models.ToPointer("name2"),
                Scope:                models.ToPointer(models.MetafieldScope{
                    Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
                    Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
                    Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
                    Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
                    PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
                }),
                DataCount:            models.ToPointer(10),
                InputType:            models.ToPointer(models.MetafieldInput_BALANCETRACKER),
            },
            models.Metafield{
                Id:                   models.ToPointer(22),
                Name:                 models.ToPointer("name2"),
                Scope:                models.ToPointer(models.MetafieldScope{
                    Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
                    Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
                    Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
                    Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
                    PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
                }),
                DataCount:            models.ToPointer(10),
                InputType:            models.ToPointer(models.MetafieldInput_BALANCETRACKER),
            },
        },
    }

}
```

