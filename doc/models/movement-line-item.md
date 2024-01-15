
# Movement Line Item

## Structure

`MovementLineItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductId` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | For Product (or "baseline") line items, this field will have a value of `0`. |
| `PricePointId` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Mrr` | `*int` | Optional | - |
| `MrrMovements` | [`[]models.MRRMovement`](../../doc/models/mrr-movement.md) | Optional | - |
| `Quantity` | `*int` | Optional | - |
| `PrevQuantity` | `*int` | Optional | - |
| `Recurring` | `*bool` | Optional | When `true`, the line item's MRR value will contribute to the `plan` breakout. When `false`, the line item contributes to the `usage` breakout. |

## Example (as JSON)

```json
{
  "product_id": 156,
  "component_id": 68,
  "price_point_id": 164,
  "name": "name6",
  "mrr": 154
}
```

