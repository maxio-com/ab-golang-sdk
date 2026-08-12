
# Update Metafield

## Structure

`UpdateMetafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentName` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected. Defaults to 'text'. |
| `Enum` | `[]string` | Optional | Only applicable when input_type is radio or dropdown. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateMetafield := models.UpdateMetafield{
        CurrentName:          models.ToPointer("current_name6"),
        Name:                 models.ToPointer("name2"),
        Scope:                models.ToPointer(models.MetafieldScope{
            Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
            Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
            Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
            Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
            PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
        }),
        InputType:            models.ToPointer(models.MetafieldInput_RADIO),
        Enum:                 []string{
            "enum8",
            "enum9",
            "enum0",
        },
    }

}
```

