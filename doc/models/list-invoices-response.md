
# List Invoices Response

## Structure

`ListInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoices` | [`[]models.Invoice`](invoice.md) | Required | - |

## Example (as JSON)

```json
{
  "invoices": [
    {
      "id": 196,
      "uid": "uid6",
      "site_id": 122,
      "customer_id": 234,
      "subscription_id": 50
    }
  ]
}
```

