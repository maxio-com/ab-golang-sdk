
# Scheduled Renewal Item Request Body Product

## Structure

`ScheduledRenewalItemRequestBodyProduct`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ItemType` | `string` | Required, Constant | Item type to add. Either Product or Component.<br><br>**Value**: `"Product"` |
| `ItemId` | `int` | Required | Product or component identifier. |
| `PricePointId` | `*int` | Optional | Price point identifier. |
| `Quantity` | `*int` | Optional | Optional quantity for the item. |
| `CustomPrice` | [`*models.ScheduledRenewalProductPricePoint`](../../doc/models/scheduled-renewal-product-price-point.md) | Optional | Custom pricing for a product within a scheduled renewal. |

## Example (as JSON)

```json
{
  "item_type": "Product",
  "item_id": 32,
  "price_point_id": 18,
  "quantity": 96,
  "custom_price": {
    "name": "name4",
    "handle": "handle0",
    "price_in_cents": "String3",
    "interval": "String3",
    "interval_unit": "day",
    "tax_included": false,
    "initial_charge_in_cents": 30,
    "expiration_interval": 52
  }
}
```

