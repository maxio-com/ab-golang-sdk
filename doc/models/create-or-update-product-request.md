
# Create or Update Product Request

## Structure

`CreateOrUpdateProductRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Product` | [`models.CreateOrUpdateProduct`](../../doc/models/create-or-update-product.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOrUpdateProductRequest := models.CreateOrUpdateProductRequest{
        Product:              models.CreateOrUpdateProduct{
            Name:                   "name0",
            Handle:                 models.ToPointer("handle6"),
            Description:            "description0",
            AccountingCode:         models.ToPointer("accounting_code6"),
            RequireCreditCard:      models.ToPointer(false),
            PriceInCents:           int64(54),
            Interval:               186,
            IntervalUnit:           models.IntervalUnit_DAY,
            TrialPriceInCents:      models.ToPointer(int64(34)),
            TrialInterval:          models.ToPointer(88),
        },
    }

}
```

