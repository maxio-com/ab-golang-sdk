
# Invoice Response

## Structure

`InvoiceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoice` | [`models.Invoice`](../../doc/models/invoice.md) | Required | - |

## Example (as JSON)

```json
{
  "invoice": {
    "issue_date": "2024-01-01",
    "due_date": "2024-01-01",
    "paid_date": "2024-01-01",
    "id": 166,
    "uid": "uid6",
    "site_id": 92,
    "customer_id": 204,
    "subscription_id": 20
  }
}
```

