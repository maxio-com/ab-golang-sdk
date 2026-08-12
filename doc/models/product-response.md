
# Product Response

## Structure

`ProductResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Product` | [`models.Product`](../../doc/models/product.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productResponse := models.ProductResponse{
        Product:              models.Product{
            Id:                         models.ToPointer(134),
            Name:                       models.ToPointer("name0"),
            Handle:                     models.NewOptional(models.ToPointer("handle6")),
            Description:                models.NewOptional(models.ToPointer("description0")),
            AccountingCode:             models.NewOptional(models.ToPointer("accounting_code6")),
        },
    }

}
```

