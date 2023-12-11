
# Bulk Update Segments Item

## Structure

`BulkUpdateSegmentsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | The ID of the segment you want to update. |
| `PricingScheme` | [`models.PricingSchemeEnum`](pricing-scheme-enum.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](create-or-update-segment-price.md) | Required | - |

## Example (as JSON)

```json
{
  "id": 180,
  "pricing_scheme": "per_unit",
  "prices": [
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

