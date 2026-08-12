
# Product Family Response

## Structure

`ProductFamilyResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductFamily` | [`*models.ProductFamily`](../../doc/models/product-family.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productFamilyResponse := models.ProductFamilyResponse{
        ProductFamily:        models.ToPointer(models.ProductFamily{
            Id:                   models.ToPointer(14),
            Name:                 models.ToPointer("name0"),
            Handle:               models.ToPointer("handle6"),
            AccountingCode:       models.NewOptional(models.ToPointer("accounting_code6")),
            Description:          models.NewOptional(models.ToPointer("description0")),
        }),
    }

}
```

