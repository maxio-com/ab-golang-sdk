
# Create Product Family

## Structure

`CreateProductFamily`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `models.Optional[string]` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Surcharging` | `*bool` | Optional | Whether surcharging applies to this product family. Defaults to `true` when omitted. Only applied on sites where surcharging is enabled. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createProductFamily := models.CreateProductFamily{
        Name:                 "name0",
        Handle:               models.NewOptional(models.ToPointer("handle6")),
        Description:          models.NewOptional(models.ToPointer("description0")),
        Surcharging:          models.ToPointer(false),
    }

}
```

