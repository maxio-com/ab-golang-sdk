
# Create Offer

## Structure

`CreateOffer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `string` | Required | - |
| `Description` | `*string` | Optional | - |
| `ProductId` | `int` | Required | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `Components` | [`[]models.CreateOfferComponent`](../../doc/models/create-offer-component.md) | Optional | - |
| `Coupons` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "name": "name4",
  "handle": "handle0",
  "description": "description4",
  "product_id": 208,
  "product_price_point_id": 132,
  "components": [
    {
      "component_id": 108,
      "price_point_id": 124,
      "starting_quantity": 84
    },
    {
      "component_id": 108,
      "price_point_id": 124,
      "starting_quantity": 84
    }
  ],
  "coupons": [
    "coupons4"
  ]
}
```

