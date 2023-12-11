
# Component Price Point

## Structure

`ComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Type` | [`*models.PricePointTypeEnum`](price-point-type-enum.md) | Optional | Price point type. We expose the following types:<br><br>1. **default**: a price point that is marked as a default price for a certain product.<br>2. **custom**: a custom price point.<br>3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one. |
| `Default` | `*bool` | Optional | Note: Refer to type attribute instead |
| `Name` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingSchemeEnum`](pricing-scheme-enum.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `ComponentId` | `*int` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `ArchivedAt` | `Optional[string]` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `UpdatedAt` | `*string` | Optional | - |
| `Prices` | [`[]models.ComponentPricePointPrice`](component-price-point-price.md) | Optional | - |
| `UseSiteExchangeRate` | `*bool` | Optional | Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.<br>**Default**: `true` |
| `SubscriptionId` | `*int` | Optional | (only used for Custom Pricing - ie. when the price point's type is `custom`) The id of the subscription that the custom price point is for. |
| `TaxIncluded` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "use_site_exchange_rate": true,
  "id": 190,
  "type": "custom",
  "default": false,
  "name": "name2",
  "pricing_scheme": "stairstep"
}
```

