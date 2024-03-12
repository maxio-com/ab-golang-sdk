
# Create or Update Segment Price

## Structure

`CreateOrUpdateSegmentPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartingQuantity` | `*int` | Optional | - |
| `EndingQuantity` | `*int` | Optional | - |
| `UnitPrice` | [`models.CreateOrUpdateSegmentPriceUnitPrice`](../../doc/models/containers/create-or-update-segment-price-unit-price.md) | Required | This is a container for one-of cases. |

## Example (as JSON)

```json
{
  "starting_quantity": 78,
  "ending_quantity": 52,
  "unit_price": "String7"
}
```

