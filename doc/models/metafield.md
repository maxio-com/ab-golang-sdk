
# Metafield

## Structure

`Metafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `DataCount` | `*int` | Optional | The amount of subscriptions this metafield has been applied to in Advanced Billing. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected. Defaults to 'text'. |
| `Enum` | [`models.Optional[models.MetafieldEnum]`](../../doc/models/containers/metafield-enum.md) | Optional | This is a container for one-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    metafield := models.Metafield{
        Id:                   models.ToPointer(242),
        Name:                 models.ToPointer("name4"),
        Scope:                models.ToPointer(models.MetafieldScope{
            Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
            Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
            Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
            Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
            PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
        }),
        DataCount:            models.ToPointer(26),
        InputType:            models.ToPointer(models.MetafieldInput_BALANCETRACKER),
    }

}
```

