
# Create or Update Product Request

## Structure

`CreateOrUpdateProductRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Product` | [`models.CreateOrUpdateProduct`](../../doc/models/create-or-update-product.md) | Required | - |

## Example (as JSON)

```json
{
  "product": {
    "name": "name0",
    "handle": "handle6",
    "description": "description0",
    "accounting_code": "accounting_code6",
    "require_credit_card": false,
    "price_in_cents": 54,
    "interval": 186,
    "interval_unit": "day",
    "auto_create_signup_page": false,
    "tax_code": "tax_code8"
  }
}
```

