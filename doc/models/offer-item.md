
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
| `CurrencyPrices` | [`[]models.CurrencyPrice`](../../doc/models/currency-price.md) | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |

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

