
# Create Product Family Request

## Structure

`CreateProductFamilyRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductFamily` | [`models.CreateProductFamily`](../../doc/models/create-product-family.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createProductFamilyRequest := models.CreateProductFamilyRequest{
        ProductFamily:        models.CreateProductFamily{
            Name:                 "name0",
            Handle:               models.NewOptional(models.ToPointer("handle6")),
            Description:          models.NewOptional(models.ToPointer("description0")),
            Surcharging:          models.ToPointer(false),
        },
    }

}
```

