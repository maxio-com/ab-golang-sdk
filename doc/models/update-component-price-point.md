
# Update Component Price Point

## Structure

`UpdateComponentPricePoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Prices` | [`[]models.UpdatePrice`](update-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "name": "name2",
  "prices": [
    {
      "id": 18,
      "ending_quantity": 38,
      "unit_price": 88,
      "_destroy": "_destroy4",
      "starting_quantity": 64
    },
    {
      "id": 18,
      "ending_quantity": 38,
      "unit_price": 88,
      "_destroy": "_destroy4",
      "starting_quantity": 64
    }
  ]
}
```

