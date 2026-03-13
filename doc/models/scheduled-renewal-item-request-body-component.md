
# Scheduled Renewal Item Request Body Component

## Structure

`ScheduledRenewalItemRequestBodyComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ItemType` | `string` | Required, Constant | Item type to add. Either Product or Component.<br><br>**Value**: `"Component"` |
| `ItemId` | `int` | Required | Product or component identifier. |
| `PricePointId` | `*int` | Optional | Price point identifier. |
| `Quantity` | `*int` | Optional | Optional quantity for the item. |
| `CustomPrice` | [`*models.ScheduledRenewalComponentCustomPrice`](../../doc/models/scheduled-renewal-component-custom-price.md) | Optional | Custom pricing for a component within a scheduled renewal. |

## Example (as JSON)

```json
{
  "item_type": "Component",
  "item_id": 108,
  "price_point_id": 122,
  "quantity": 212,
  "custom_price": {
    "tax_included": false,
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      },
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ]
  }
}
```

