
# Component Cost Data

## Structure

`ComponentCostData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentCodeId` | `models.Optional[int]` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `Quantity` | `*string` | Optional | - |
| `Amount` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Tiers` | [`[]models.ComponentCostDataRateTier`](../../doc/models/component-cost-data-rate-tier.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentCostData := models.ComponentCostData{
        ComponentCodeId:      models.NewOptional(models.ToPointer(40)),
        PricePointId:         models.ToPointer(210),
        ProductId:            models.ToPointer(18),
        Quantity:             models.ToPointer("quantity4"),
        Amount:               models.ToPointer("amount0"),
    }

}
```

