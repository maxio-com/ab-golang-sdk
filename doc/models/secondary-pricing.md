
# Secondary Pricing

Hybrid Pricing only. The secondary pricing model that takes over once usage passes the primary model's included threshold. Applies to Quantity-Based and Metered Components on Sites with Hybrid Pricing enabled; ignored otherwise. Its `prices` brackets start at `1`, not at the primary model's highest `ending_quantity`. See [Hybrid Pricing](http://localhost:8080/go).

## Structure

`SecondaryPricing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | Required, unless you send the top-level `secondary_pricing_scheme` instead (Create and Update Component Price Point only). A block with neither returns `422` (`Secondary pricing scheme must be defined`). |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    secondaryPricing := models.SecondaryPricing{
        PricingScheme:        models.ToPointer(models.PricingScheme_STAIRSTEP),
        Prices:               []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
        },
    }

}
```

