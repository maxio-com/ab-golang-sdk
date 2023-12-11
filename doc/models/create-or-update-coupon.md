
# Create or Update Coupon

## Structure

`CreateOrUpdateCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coupon` | `*interface{}` | Optional | - |
| `RestrictedProducts` | `map[string]bool` | Optional | An object where the keys are product_ids and the values are booleans indicating if the coupon should be applicable to the product |
| `RestrictedComponents` | `map[string]bool` | Optional | An object where the keys are component_ids and the values are booleans indicating if the coupon should be applicable to the component |

## Example (as JSON)

```json
{
  "coupon": {
    "key1": "val1",
    "key2": "val2"
  },
  "restricted_products": {
    "key0": true
  },
  "restricted_components": {
    "key0": true
  }
}
```

