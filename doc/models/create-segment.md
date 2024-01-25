
# Create Segment

## Structure

`CreateSegment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SegmentProperty1Value` | `*interface{}` | Optional | A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric. |
| `SegmentProperty2Value` | `*interface{}` | Optional | A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric. |
| `SegmentProperty3Value` | `*interface{}` | Optional | A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric. |
| `SegmentProperty4Value` | `*interface{}` | Optional | A value that will occur in your events that you want to bill upon. The type of the value depends on the property type in the related event based billing metric. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "segment_property_1_value": {
    "key1": "val1",
    "key2": "val2"
  },
  "segment_property_2_value": {
    "key1": "val1",
    "key2": "val2"
  },
  "segment_property_3_value": {
    "key1": "val1",
    "key2": "val2"
  },
  "segment_property_4_value": {
    "key1": "val1",
    "key2": "val2"
  },
  "pricing_scheme": "per_unit",
  "prices": [
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ]
}
```

