
# Invoice Line Item Component Cost Data

## Structure

`InvoiceLineItemComponentCostData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Rates` | [`[]models.ComponentCostData`](../../doc/models/component-cost-data.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceLineItemComponentCostData := models.InvoiceLineItemComponentCostData{
        Rates:                []models.ComponentCostData{
            models.ComponentCostData{
                ComponentCodeId:      models.NewOptional(models.ToPointer(116)),
                PricePointId:         models.ToPointer(226),
                ProductId:            models.ToPointer(94),
                Quantity:             models.ToPointer("quantity0"),
                Amount:               models.ToPointer("amount6"),
            },
            models.ComponentCostData{
                ComponentCodeId:      models.NewOptional(models.ToPointer(116)),
                PricePointId:         models.ToPointer(226),
                ProductId:            models.ToPointer(94),
                Quantity:             models.ToPointer("quantity0"),
                Amount:               models.ToPointer("amount6"),
            },
        },
    }

}
```

