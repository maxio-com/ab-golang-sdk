
# Create Invoice Request

## Structure

`CreateInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoice` | [`models.CreateInvoice`](../../doc/models/create-invoice.md) | Required | - |

## Example (as JSON)

```json
{
  "invoice": {
    "issue_date": "2024-01-01",
    "status": "draft",
    "line_items": [
      {
        "title": "title4",
        "quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        },
        "taxable": false,
        "tax_code": "tax_code6"
      },
      {
        "title": "title4",
        "quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        },
        "taxable": false,
        "tax_code": "tax_code6"
      },
      {
        "title": "title4",
        "quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        },
        "taxable": false,
        "tax_code": "tax_code6"
      }
    ],
    "net_terms": 144,
    "payment_instructions": "payment_instructions6",
    "memo": "memo0"
  }
}
```

