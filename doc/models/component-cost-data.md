
# Component Cost Data

## Structure

`ComponentCostData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentCodeId` | `Optional[int]` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `Quantity` | `*string` | Optional | - |
| `Amount` | `*string` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Tiers` | [`[]models.ComponentCostDataRateTier`](../../doc/models/component-cost-data-rate-tier.md) | Optional | - |

## Example (as JSON)

```json
{
  "component_code_id": 16,
  "price_point_id": 186,
  "product_id": 250,
  "quantity": "quantity8",
  "amount": "amount4"
}
```

