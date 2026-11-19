
# Update Secondary Pricing

Hybrid Pricing only. Changes the secondary pricing model of an existing hybrid price point. Send `pricing_scheme` (or the top-level `secondary_pricing_scheme`) every time you send this block; a block without a scheme returns `422`. Brackets sent without an `id` are added to the existing ones rather than replacing them, so to change a bracket set, update existing brackets by `id` and remove the ones you no longer want with `_destroy: true`. See [Hybrid Pricing](http://localhost:8080/go).

## Structure

`UpdateSecondaryPricing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.UpdatePrice`](../../doc/models/update-price.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSecondaryPricing := models.UpdateSecondaryPricing{
        PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
        Prices:               []models.UpdatePrice{
            models.UpdatePrice{
                Id:                   models.ToPointer(18),
                EndingQuantity:       models.ToPointer(models.UpdatePriceEndingQuantityContainer.FromNumber(40)),
                UnitPrice:            models.ToPointer(models.UpdatePriceUnitPriceContainer.FromPrecision(float64(23.26))),
                Destroy:              models.ToPointer(false),
                StartingQuantity:     models.ToPointer(models.UpdatePriceStartingQuantityContainer.FromNumber(242)),
            },
        },
    }

}
```

