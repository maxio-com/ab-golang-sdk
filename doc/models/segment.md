
# Segment

## Structure

`Segment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `ComponentId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |
| `EventBasedBillingMetricId` | `*int` | Optional | - |
| `PricingScheme` | [`*models.PricingScheme`](../../doc/models/pricing-scheme.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `SegmentProperty1Value` | [`*models.SegmentSegmentProperty1Value`](../../doc/models/containers/segment-segment-property-1-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty2Value` | [`*models.SegmentSegmentProperty2Value`](../../doc/models/containers/segment-segment-property-2-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty3Value` | [`*models.SegmentSegmentProperty3Value`](../../doc/models/containers/segment-segment-property-3-value.md) | Optional | This is a container for one-of cases. |
| `SegmentProperty4Value` | [`*models.SegmentSegmentProperty4Value`](../../doc/models/containers/segment-segment-property-4-value.md) | Optional | This is a container for one-of cases. |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `Prices` | [`[]models.SegmentPrice`](../../doc/models/segment-price.md) | Optional | **Constraints**: *Minimum Items*: `1` |

## Example (as JSON)

```json
{
  "id": 6,
  "component_id": 116,
  "price_point_id": 140,
  "event_based_billing_metric_id": 200,
  "pricing_scheme": "stairstep"
}
```

