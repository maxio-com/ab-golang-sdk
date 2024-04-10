
# Create Segment

## Structure

`CreateSegment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SegmentProperty1Value` | [`*models.CreateSegmentSegmentProperty1Value`](../../doc/models/containers/create-segment-segment-property-1-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty2Value` | [`*models.CreateSegmentSegmentProperty2Value`](../../doc/models/containers/create-segment-segment-property-2-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty3Value` | [`*models.CreateSegmentSegmentProperty3Value`](../../doc/models/containers/create-segment-segment-property-3-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty4Value` | [`*models.CreateSegmentSegmentProperty4Value`](../../doc/models/containers/create-segment-segment-property-4-value.md) | Optional | This is a container for one-of cases. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Optional | - |

## Example (as JSON)

```json
{
  "segment_property_1_value": "String9",
  "segment_property_2_value": "String1",
  "segment_property_3_value": "String3",
  "segment_property_4_value": "String3",
  "pricing_scheme": "per_unit",
  "prices": [
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": "String3"
    },
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": "String3"
    },
    {
      "starting_quantity": 64,
      "ending_quantity": 38,
      "unit_price": "String3"
    }
  ]
}
```

