
# Create Prepaid Usage Component Price Point

## Structure

`CreatePrepaidUsageComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `*string` | Optional | - |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Required | - |
| `OveragePricing` | [`models.OveragePricing`](../../doc/models/overage-pricing.md) | Required | - |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.<br><br>**Default**: `true` |
| `RolloverPrepaidRemainder` | `*bool` | Optional | (only for prepaid usage components) Boolean which controls whether or not remaining units should be rolled over to the next period. |
| `RenewPrepaidAllocation` | `*bool` | Optional | (only for prepaid usage components) Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period. |
| `ExpirationInterval` | `*float64` | Optional | (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | (only for prepaid usage components where rollover_prepaid_remainder is true) A string representing the expiration interval unit for this component, either month or day. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPrepaidUsageComponentPricePoint := models.CreatePrepaidUsageComponentPricePoint{
        Name:                     "name2",
        Handle:                   models.ToPointer("handle8"),
        PricingScheme:            models.PricingScheme_PERUNIT,
        Prices:                   []models.Price{
            models.Price{
                StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
            },
        },
        OveragePricing:           models.OveragePricing{
            PricingScheme:        models.PricingScheme_STAIRSTEP,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
            },
        },
        UseSiteExchangeRate:      models.ToPointer(true),
        RolloverPrepaidRemainder: models.ToPointer(false),
        RenewPrepaidAllocation:   models.ToPointer(false),
        ExpirationInterval:       models.ToPointer(float64(165.7)),
    }

}
```

