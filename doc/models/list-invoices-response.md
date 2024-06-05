
# List Invoices Response

## Structure

`ListInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoices` | [`[]models.Invoice`](../../doc/models/invoice.md) | Required | - |

## Example (as JSON)

```json
{
  "invoices": [
    {
      "issue_date": "2024-01-01",
      "due_date": "2024-01-01",
      "paid_date": "2024-01-01",
      "id": 196,
      "uid": "uid6",
      "site_id": 122,
      "customer_id": 234,
      "subscription_id": 50
    }
  ]
}
```

