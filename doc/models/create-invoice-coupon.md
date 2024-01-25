
# Create Invoice Coupon

## Structure

`CreateInvoiceCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | - |
| `Percentage` | `*interface{}` | Optional | - |
| `Amount` | `*interface{}` | Optional | - |
| `Description` | `*string` | Optional | **Constraints**: *Maximum Length*: `255` |
| `ProductFamilyId` | `*interface{}` | Optional | - |
| `CompoundingStrategy` | [`*models.CompoundingStrategy`](../../doc/models/compounding-strategy.md) | Optional | - |

## Example (as JSON)

```json
{
  "percentage": {
    "key1": "val1",
    "key2": "val2"
  },
  "code": "code4",
  "amount": {
    "key1": "val1",
    "key2": "val2"
  },
  "description": "description4",
  "product_family_id": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

