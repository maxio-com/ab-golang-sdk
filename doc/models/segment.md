
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
| `PricingScheme` | [`*models.PricingSchemeEnum`](pricing-scheme-enum.md) | Optional | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `SegmentProperty1Value` | `*interface{}` | Optional | - |
| `SegmentProperty2Value` | `*interface{}` | Optional | - |
| `SegmentProperty3Value` | `*interface{}` | Optional | - |
| `SegmentProperty4Value` | `*interface{}` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `UpdatedAt` | `*string` | Optional | - |
| `Prices` | [`[]models.SegmentPrice`](segment-price.md) | Optional | **Constraints**: *Minimum Items*: `1` |

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

