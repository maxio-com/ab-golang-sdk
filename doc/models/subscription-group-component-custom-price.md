
# Subscription Group Component Custom Price

Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`.

## Structure

`SubscriptionGroupComponentCustomPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`*models.PricingSchemeEnum`](pricing-scheme-enum.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.Price`](price.md) | Optional | - |
| `OveragePricing` | [`[]models.ComponentCustomPrice`](component-custom-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "pricing_scheme": "per_unit",
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
  ],
  "overage_pricing": [
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
  ]
}
```

