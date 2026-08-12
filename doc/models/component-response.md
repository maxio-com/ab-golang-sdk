
# Component Response

## Structure

`ComponentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Component` | [`models.Component`](../../doc/models/component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentResponse := models.ComponentResponse{
        Component:            models.Component{
            Id:                        models.ToPointer(80),
            Name:                      models.ToPointer("name8"),
            Handle:                    models.NewOptional(models.ToPointer("handle4")),
            PricingScheme:             models.NewOptional(models.ToPointer(models.PricingScheme_PERUNIT)),
            UnitName:                  models.ToPointer("unit_name0"),
            ItemCategory:              models.NewOptional(models.ToPointer(models.ItemCategory_ENUMBUSINESSSOFTWARE)),
        },
    }

}
```

