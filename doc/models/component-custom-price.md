
# Component Custom Price

Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.

## Structure

`ComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`*models.PricingSchemeEnum`](pricing-scheme-enum.md) | Optional | Omit for On/Off components |
| `Prices` | [`[]models.Price`](price.md) | Optional | On/off components only need one price bracket starting at 1 |

## Example (as JSON)

```json
{
  "pricing_scheme": "stairstep",
  "prices": [
    {
      "starting_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "ending_quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ]
}
```

