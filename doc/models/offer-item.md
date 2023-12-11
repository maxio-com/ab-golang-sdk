
# Offer Item

## Structure

`OfferItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `StartingQuantity` | `*string` | Optional | - |
| `Editable` | `*bool` | Optional | - |
| `ComponentUnitPrice` | `*string` | Optional | - |
| `ComponentName` | `*string` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `CurrencyPrices` | [`[]models.CurrencyPrice`](currency-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "component_id": 216,
  "price_point_id": 16,
  "starting_quantity": "starting_quantity0",
  "editable": false,
  "component_unit_price": "component_unit_price8"
}
```

