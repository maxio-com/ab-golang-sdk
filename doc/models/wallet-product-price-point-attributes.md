
# Wallet Product Price Point Attributes

A price point to create for the Wallet Product, in the same call that creates or updates it.

## Structure

`WalletProductPricePointAttributes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | The price point's name. Required. Omitting it fails validation with a message that reads `Name: cannot be blank.`, not naming the price point. |
| `Handle` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `UnitPrice` | `models.Optional[string]` | Optional | Required when `pricing_scheme` is `per_unit`. |
| `TaxIncluded` | `*bool` | Optional | - |
| `Interval` | `*int` | Optional | Requires Multifrequency to be enabled on the site. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | - |
| `UseSiteExchangeRate` | `*bool` | Optional | - |
| `Prices` | [`[]models.Price1`](../../doc/models/price-1.md) | Optional | Required for every `pricing_scheme` other than `per_unit`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductPricePointAttributes := models.WalletProductPricePointAttributes{
        Name:                 "name8",
        Handle:               models.ToPointer("handle4"),
        PricingScheme:        models.ToPointer(models.PricingScheme_PERUNIT),
        UnitPrice:            models.NewOptional(models.ToPointer("unit_price6")),
        TaxIncluded:          models.ToPointer(false),
        Interval:             models.ToPointer(120),
    }

}
```

