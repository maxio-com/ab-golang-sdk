
# Create Component Price Point

## Structure

`CreateComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `*string` | Optional | - |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Required | - |
| `SecondaryPricing` | [`*models.SecondaryPricing`](../../doc/models/secondary-pricing.md) | Optional | Hybrid Pricing only. The secondary pricing model that takes over once usage passes the primary model's included threshold. Applies to Quantity-Based and Metered Components on Sites with Hybrid Pricing enabled; ignored otherwise. Its `prices` brackets start at `1`, not at the primary model's highest `ending_quantity`. See [Hybrid Pricing](http://localhost:8080/go). |
| `SecondaryPricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | Hybrid Pricing only. Alternative to `secondary_pricing.pricing_scheme`; when both are sent, `secondary_pricing.pricing_scheme` is used. Setting not supported when creating price points in bulk. |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site. Setting not supported when creating price points in bulk.<br><br>**Default**: `true` |
| `TaxIncluded` | `*bool` | Optional | Whether or not the price point includes tax. Setting not supported when creating price points in bulk. |
| `Interval` | `*int` | Optional | The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this price point, either month or day. This property is only available for sites with Multifrequency enabled. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createComponentPricePoint := models.CreateComponentPricePoint{
        Name:                   "name4",
        Handle:                 models.ToPointer("handle0"),
        PricingScheme:          models.PricingScheme_PERUNIT,
        Prices:                 []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
        },
        SecondaryPricing:       models.ToPointer(models.SecondaryPricing{
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
            },
        }),
        SecondaryPricingScheme: models.ToPointer(models.PricingScheme_STAIRSTEP),
        UseSiteExchangeRate:    models.ToPointer(true),
        TaxIncluded:            models.ToPointer(false),
    }

}
```

