
# Update Component Request

## Structure

`UpdateComponentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Component` | [`models.UpdateComponent`](../../doc/models/update-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateComponentRequest := models.UpdateComponentRequest{
        Component:            models.UpdateComponent{
            Handle:               models.ToPointer("handle4"),
            Name:                 models.ToPointer("name8"),
            Description:          models.NewOptional(models.ToPointer("description2")),
            AccountingCode:       models.NewOptional(models.ToPointer("accounting_code4")),
            Taxable:              models.ToPointer(false),
            ItemCategory:         models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
        },
    }

}
```

