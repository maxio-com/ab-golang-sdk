
# Prepaid Usage Component

## Structure

`PrepaidUsageComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | A name for this component that is suitable for showing customers and displaying on billing statements, e.g., "Minutes". |
| `UnitName` | `string` | Required | The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. e.g., “message”, which may then be shown as “5 messages” on a subscription’s component line-item |
| `Description` | `*string` | Optional | A description for the component that will be displayed to the user on the hosted signup page. |
| `Handle` | `*string` | Optional | A unique identifier for your use that can be used to retrieve this component in subsequent requests. Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.<br><br>**Constraints**: *Pattern*: `^[a-z0-9][a-z0-9\-_:.]*$` |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](../../doc/models/price.md) | Optional | (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261149166733-Component-Pricing-Schemes#price-bracket-rules) for an overview of how price brackets work for different pricing schemes. |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `PricePoints` | [`[]models.CreatePrepaidUsageComponentPricePoint`](../../doc/models/create-prepaid-usage-component-price-point.md) | Optional | - |
| `UnitPrice` | [`*models.PrepaidUsageComponentUnitPrice`](../../doc/models/containers/prepaid-usage-component-unit-price.md) | Optional | This is a container for one-of cases. |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `HideDateRangeOnInvoice` | `*bool` | Optional | (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices. |
| `OveragePricing` | [`models.OveragePricing`](../../doc/models/overage-pricing.md) | Required | - |
| `RolloverPrepaidRemainder` | `*bool` | Optional | Boolean which controls whether or not remaining units should be rolled over to the next period. |
| `RenewPrepaidAllocation` | `*bool` | Optional | Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period. |
| `ExpirationInterval` | `*float64` | Optional | (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | - |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `PublicSignupPageIds` | `[]int` | Optional | - |
| `UnspscCode` | `models.Optional[string]` | Optional | (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this component instead of the default derived from item_category. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidUsageComponent := models.PrepaidUsageComponent{
        Name:                      "name2",
        UnitName:                  "unit_name4",
        Description:               models.ToPointer("description8"),
        Handle:                    models.ToPointer("handle8"),
        Taxable:                   models.ToPointer(false),
        PricingScheme:             models.PricingScheme_PERUNIT,
        Prices:                    []models.Price{
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
        UpgradeCharge:             models.NewOptional(models.ToPointer(models.CreditType_NONE)),
        OveragePricing:            models.OveragePricing{
            PricingScheme:        models.PricingScheme_STAIRSTEP,
            Prices:               []models.Price{
                models.Price{
                    StartingQuantity:     models.PriceStartingQuantityContainer.FromNumber(242),
                    EndingQuantity:       models.NewOptional(models.ToPointer(models.PriceEndingQuantityContainer.FromNumber(40))),
                    UnitPrice:            models.PriceUnitPriceContainer.FromPrecision(float64(23.26)),
                },
            },
        },
    }

}
```

