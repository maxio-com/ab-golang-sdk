
# Create Metafield

## Structure

`CreateMetafield`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Scope` | [`*models.MetafieldScope`](../../doc/models/metafield-scope.md) | Optional | Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings. |
| `InputType` | [`*models.MetafieldInput`](../../doc/models/metafield-input.md) | Optional | Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected. Defaults to 'text'. |
| `Enum` | `[]string` | Optional | Only applicable when input_type is radio or dropdown. Empty strings will not be submitted. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createMetafield := models.CreateMetafield{
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
    }

}
```

