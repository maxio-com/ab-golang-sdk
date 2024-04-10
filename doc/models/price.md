
# Price

## Structure

`Price`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | [`models.PriceStartingQuantity`](../../doc/models/containers/price-starting-quantity.md) | Required | This is a container for one-of cases. |
| `EndingQuantity` | [`models.Optional[models.PriceEndingQuantity]`](../../doc/models/containers/price-ending-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitPrice` | [`models.PriceUnitPrice`](../../doc/models/containers/price-unit-price.md) | Required | This is a container for one-of cases. |

## Example (as JSON)

```json
{
  "starting_quantity": 40,
  "ending_quantity": 14,
  "unit_price": 125.12
}
```

